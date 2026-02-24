package model

import (
	"time"

	"gorm.io/gorm"
)

// ========================================
// 数据库模型
// ========================================

// User 用户模型
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"size:50;uniqueIndex" json:"username"`
	Password  string         `gorm:"size:100" json:"-"`
	Email     string         `gorm:"size:100" json:"email"`
	Nickname  string         `gorm:"size:50" json:"nickname"`
	Avatar    string         `gorm:"size:500" json:"avatar"`
	Bio       string         `gorm:"size:500" json:"bio"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Article 文章模型
type Article struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"size:200" json:"title"`
	Content   string         `gorm:"type:text" json:"content"`
	Summary   string         `gorm:"size:500" json:"summary"`
	Cover     string         `gorm:"size:500" json:"cover"`
	Category  string         `gorm:"size:50;index" json:"category"`
	Tags      string         `gorm:"size:200" json:"tags"`
	Views     int            `gorm:"default:0" json:"views"`
	Published bool           `gorm:"default:false;index" json:"published"`
	AuthorID  uint           `gorm:"index" json:"author_id"`
	Author    *User          `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Portfolio 作品集模型
type Portfolio struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"size:200" json:"title"`
	Description string         `gorm:"size:1000" json:"description"`
	Content     string         `gorm:"type:text" json:"content"`
	ImageURL    string         `gorm:"size:500" json:"image_url"`
	ProjectURL  string         `gorm:"size:500" json:"project_url"`
	GithubURL   string         `gorm:"size:500" json:"github_url"`
	TechStack   string         `gorm:"size:500" json:"tech_stack"`
	SortOrder   int            `gorm:"default:0" json:"sort_order"`
	Published   bool           `gorm:"default:false;index" json:"published"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// AutoMigrate 自动迁移数据库
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Article{}, &Portfolio{})
}

// ========================================
// 请求结构
// ========================================

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateArticleRequest 创建文章请求
type CreateArticleRequest struct {
	Title     string `json:"title" binding:"required,max=200"`
	Content   string `json:"content" binding:"required"`
	Summary   string `json:"summary" binding:"max=500"`
	Cover     string `json:"cover" binding:"max=500"`
	Category  string `json:"category" binding:"max=50"`
	Tags      string `json:"tags" binding:"max=200"`
	Published bool   `json:"published"`
}

// UpdateArticleRequest 更新文章请求
type UpdateArticleRequest struct {
	Title     *string `json:"title" binding:"omitempty,max=200"`
	Content   *string `json:"content"`
	Summary   *string `json:"summary" binding:"omitempty,max=500"`
	Cover     *string `json:"cover" binding:"omitempty,max=500"`
	Category  *string `json:"category" binding:"omitempty,max=50"`
	Tags      *string `json:"tags" binding:"omitempty,max=200"`
	Published *bool   `json:"published"`
}

// CreatePortfolioRequest 创建作品集请求
type CreatePortfolioRequest struct {
	Title       string `json:"title" binding:"required,max=200"`
	Description string `json:"description" binding:"max=1000"`
	Content     string `json:"content"`
	ImageURL    string `json:"image_url" binding:"max=500"`
	ProjectURL  string `json:"project_url" binding:"max=500"`
	GithubURL   string `json:"github_url" binding:"max=500"`
	TechStack   string `json:"tech_stack" binding:"max=500"`
	SortOrder   int    `json:"sort_order"`
	Published   bool   `json:"published"`
}

// UpdatePortfolioRequest 更新作品集请求
type UpdatePortfolioRequest struct {
	Title       *string `json:"title" binding:"omitempty,max=200"`
	Description *string `json:"description" binding:"omitempty,max=1000"`
	Content     *string `json:"content"`
	ImageURL    *string `json:"image_url" binding:"omitempty,max=500"`
	ProjectURL  *string `json:"project_url" binding:"omitempty,max=500"`
	GithubURL   *string `json:"github_url" binding:"omitempty,max=500"`
	TechStack   *string `json:"tech_stack" binding:"omitempty,max=500"`
	SortOrder   *int    `json:"sort_order"`
	Published   *bool   `json:"published"`
}

// UpdateProfileRequest 更新用户资料请求
type UpdateProfileRequest struct {
	Nickname *string `json:"nickname" binding:"omitempty,max=50"`
	Email    *string `json:"email" binding:"omitempty,max=100,email"`
	Avatar   *string `json:"avatar" binding:"omitempty,max=500"`
	Bio      *string `json:"bio" binding:"omitempty,max=500"`
}

// ========================================
// 响应结构
// ========================================

// LoginResponse 登录响应
type LoginResponse struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

// PaginationQuery 分页查询
type PaginationQuery struct {
	Page     int    `form:"page" binding:"min=1"`
	PageSize int    `form:"page_size" binding:"min=1,max=100"`
	Category string `form:"category"`
	Search   string `form:"search"`
}

// PaginatedResponse 分页响应
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}
