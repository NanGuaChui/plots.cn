package handler

import (
	"net/http"
	"strings"
	"time"

	"plots/internal/config"
	"plots/internal/model"
	"plots/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Handler 处理器
type Handler struct {
	repos *repository.Repositories
	cfg   *config.Config
}

// NewHandler 创建处理器
func NewHandler(repos *repository.Repositories, cfg *config.Config) *Handler {
	return &Handler{repos: repos, cfg: cfg}
}

// Response 统一响应格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{Code: 0, Message: "success", Data: data})
}

// Error 错误响应
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{Code: code, Message: message})
}

// Login 登录
func (h *Handler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	user, err := h.repos.User.GetByUsername(req.Username)
	if err != nil {
		Error(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		Error(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	token, err := h.generateToken(user.ID)
	if err != nil {
		Error(c, http.StatusInternalServerError, "生成token失败")
		return
	}

	Success(c, model.LoginResponse{Token: token, User: user})
}

// generateToken 生成 JWT token
func (h *Handler) generateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(h.cfg.JWT.Expire).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(h.cfg.JWT.Secret))
}

// AuthMiddleware 认证中间件
func (h *Handler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			Error(c, http.StatusUnauthorized, "未授权")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			Error(c, http.StatusUnauthorized, "无效的token格式")
			c.Abort()
			return
		}

		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(h.cfg.JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			Error(c, http.StatusUnauthorized, "无效的token")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			Error(c, http.StatusUnauthorized, "无效的token")
			c.Abort()
			return
		}

		userID := uint(claims["user_id"].(float64))
		c.Set("user_id", userID)
		c.Next()
	}
}

// CheckInit 检查是否已初始化管理员
func (h *Handler) CheckInit(c *gin.Context) {
	count, err := h.repos.User.Count()
	if err != nil {
		Error(c, http.StatusInternalServerError, "检查失败")
		return
	}
	Success(c, gin.H{"initialized": count > 0})
}

// InitAdmin 初始化管理员账号（仅当没有任何用户时可用）
func (h *Handler) InitAdmin(c *gin.Context) {
	// 检查是否已有用户
	count, err := h.repos.User.Count()
	if err != nil {
		Error(c, http.StatusInternalServerError, "检查失败")
		return
	}
	if count > 0 {
		Error(c, http.StatusForbidden, "管理员已存在，无法重复初始化")
		return
	}

	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	// 验证用户名和密码长度
	if len(req.Username) < 3 || len(req.Password) < 6 {
		Error(c, http.StatusBadRequest, "用户名至少3位，密码至少6位")
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		Error(c, http.StatusInternalServerError, "密码加密失败")
		return
	}

	// 创建管理员
	admin := &model.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Nickname: "管理员",
	}

	if err := h.repos.User.Create(admin); err != nil {
		Error(c, http.StatusInternalServerError, "创建管理员失败")
		return
	}

	// 生成 token
	token, err := h.generateToken(admin.ID)
	if err != nil {
		Error(c, http.StatusInternalServerError, "生成token失败")
		return
	}

	Success(c, model.LoginResponse{Token: token, User: admin})
}

// GetProfile 获取用户资料
func (h *Handler) GetProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	user, err := h.repos.User.GetByID(userID)
	if err != nil {
		Error(c, http.StatusNotFound, "用户不存在")
		return
	}
	Success(c, user)
}

// UpdateProfile 更新用户资料
func (h *Handler) UpdateProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	var req model.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	user, err := h.repos.User.GetByID(userID)
	if err != nil {
		Error(c, http.StatusNotFound, "用户不存在")
		return
	}

	if req.Nickname != nil {
		user.Nickname = *req.Nickname
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Avatar != nil {
		user.Avatar = *req.Avatar
	}
	if req.Bio != nil {
		user.Bio = *req.Bio
	}

	if err := h.repos.User.Update(user); err != nil {
		Error(c, http.StatusInternalServerError, "更新失败")
		return
	}

	Success(c, user)
}
