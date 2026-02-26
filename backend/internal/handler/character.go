package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"plots/internal/middleware"
	"plots/internal/model"
)

// CreateCharacterRequest 创建角色请求
type CreateCharacterRequest struct {
	Name      string `json:"name" binding:"required,min=1,max=50"`
	SlotIndex int    `json:"slot_index" binding:"min=0,max=3"`
}

// GetCharacters 获取当前用户所有角色
// GET /api/characters
func GetCharacters(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	var characters []model.Character
	if err := model.DB.Where("user_id = ?", userID).Order("slot_index").Find(&characters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取角色列表失败"})
		return
	}

	// 转换为包含衍生属性的角色列表
	var result []model.CharacterWithStats
	for _, char := range characters {
		result = append(result, char.ToCharacterWithStats(nil))
	}

	c.JSON(http.StatusOK, result)
}

// CreateCharacter 创建新角色
// POST /api/characters
func CreateCharacter(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	var req CreateCharacterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误: " + err.Error()})
		return
	}

	// 检查用户角色数量（最多4个）
	var count int64
	model.DB.Model(&model.Character{}).Where("user_id = ?", userID).Count(&count)
	if count >= 4 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "角色数量已达上限（最多4个）"})
		return
	}

	// 检查槽位是否已被占用
	var existingChar model.Character
	if err := model.DB.Where("user_id = ? AND slot_index = ?", userID, req.SlotIndex).First(&existingChar).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该槽位已被占用"})
		return
	}

	// 确定角色类型：槽位0为主号，其他为铁人号
	var charType model.CharacterType
	if req.SlotIndex == 0 {
		charType = model.CharacterTypeMain
	} else {
		charType = model.CharacterTypeIronman
	}

	// 创建角色
	character := model.Character{
		UserID:       userID,
		Name:         req.Name,
		SlotIndex:    req.SlotIndex,
		Type:         charType,
		Strength:     10,
		Agility:      10,
		Endurance:    10,
		Intelligence: 10,
		Luck:         10,
		CombatLevel:  1,
		CombatExp:    0,
	}

	if err := model.DB.Create(&character).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建角色失败"})
		return
	}

	c.JSON(http.StatusOK, character.ToCharacterWithStats(nil))
}

// GetCharacter 获取角色详情
// GET /api/characters/:id
func GetCharacter(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	characterID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的角色ID"})
		return
	}

	var character model.Character
	if err := model.DB.First(&character, characterID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "角色不存在"})
		return
	}

	// 验证角色属于当前用户
	if character.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权访问该角色"})
		return
	}

	c.JSON(http.StatusOK, character.ToCharacterWithStats(nil))
}

// DeleteCharacter 删除角色
// DELETE /api/characters/:id
func DeleteCharacter(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	characterID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的角色ID"})
		return
	}

	var character model.Character
	if err := model.DB.First(&character, characterID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "角色不存在"})
		return
	}

	// 验证角色属于当前用户
	if character.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权删除该角色"})
		return
	}

	// 删除角色
	if err := model.DB.Delete(&character).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除角色失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "角色已删除"})
}

// GetCharacterStats 获取角色属性（包含装备加成）
// GET /api/characters/:id/stats
func GetCharacterStats(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	characterID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的角色ID"})
		return
	}

	var character model.Character
	if err := model.DB.First(&character, characterID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "角色不存在"})
		return
	}

	// 验证角色属于当前用户
	if character.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权访问该角色"})
		return
	}

	// TODO: 后续实现装备系统后，这里需要获取装备加成
	derivedStats := model.CalculateDerivedStats(&character, nil)

	c.JSON(http.StatusOK, gin.H{
		"character":     character,
		"derived_stats": derivedStats,
	})
}
