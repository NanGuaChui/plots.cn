package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"plots/internal/model"
)

// GetItems 获取所有物品定义
// GET /api/items
func GetItems(c *gin.Context) {
	var items []model.Item

	// 查询所有物品，预加载标签
	if err := model.DB.Preload("Tags").Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取物品列表失败"})
		return
	}

	c.JSON(http.StatusOK, items)
}

// GetItemByCode 获取单个物品详情
// GET /api/items/:code
func GetItemByCode(c *gin.Context) {
	code := c.Param("code")

	var item model.Item
	if err := model.DB.Preload("Tags").Where("code = ?", code).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "物品不存在"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// GetItemTags 获取所有物品标签
// GET /api/item-tags
func GetItemTags(c *gin.Context) {
	var tags []model.ItemTag

	if err := model.DB.Find(&tags).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取标签列表失败"})
		return
	}

	c.JSON(http.StatusOK, tags)
}

// GetItemsByTag 根据标签获取物品
// GET /api/items/tag/:tagCode
func GetItemsByTag(c *gin.Context) {
	tagCode := c.Param("tagCode")

	// 先查找标签
	var tag model.ItemTag
	if err := model.DB.Where("code = ?", tagCode).First(&tag).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "标签不存在"})
		return
	}

	// 查询关联此标签的物品
	var items []model.Item
	if err := model.DB.Preload("Tags").
		Joins("JOIN item_item_tags ON item_item_tags.item_id = items.id").
		Where("item_item_tags.item_tag_id = ?", tag.ID).
		Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取物品列表失败"})
		return
	}

	c.JSON(http.StatusOK, items)
}
