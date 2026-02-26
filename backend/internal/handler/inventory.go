package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"plots/internal/middleware"
	"plots/internal/model"
)

// GetInventory 获取角色仓库
// GET /api/characters/:id/inventory
func GetInventory(c *gin.Context) {
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

	// 验证角色属于当前用户
	var character model.Character
	if err := model.DB.First(&character, characterID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "角色不存在"})
		return
	}
	if character.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权访问该角色"})
		return
	}

	// 获取仓库物品，预加载物品信息和标签
	var inventoryItems []model.InventoryItem
	if err := model.DB.Preload("Item").Preload("Item.Tags").
		Where("character_id = ?", characterID).
		Find(&inventoryItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取仓库失败"})
		return
	}

	c.JSON(http.StatusOK, inventoryItems)
}

// AddInventoryItemRequest 添加物品请求
type AddInventoryItemRequest struct {
	ItemCode string `json:"item_code" binding:"required"` // 物品代码
	Quantity int    `json:"quantity" binding:"required,min=1"`
}

// AddInventoryItem 添加物品到仓库（测试用）
// POST /api/characters/:id/inventory/add
func AddInventoryItem(c *gin.Context) {
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

	// 验证角色属于当前用户
	var character model.Character
	if err := model.DB.First(&character, characterID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "角色不存在"})
		return
	}
	if character.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权访问该角色"})
		return
	}

	var req AddInventoryItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误: " + err.Error()})
		return
	}

	// 查找物品
	var item model.Item
	if err := model.DB.Where("code = ?", req.ItemCode).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "物品不存在"})
		return
	}

	// 添加到仓库
	inventoryItem, err := AddItemToInventory(uint(characterID), item.ID, req.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "添加物品失败: " + err.Error()})
		return
	}

	// 重新加载物品信息
	model.DB.Preload("Item").Preload("Item.Tags").First(&inventoryItem, inventoryItem.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "物品添加成功",
		"item":    inventoryItem,
	})
}

// AddItemToInventory 添加物品到仓库（内部函数）
func AddItemToInventory(characterID uint, itemID uint, quantity int) (*model.InventoryItem, error) {
	// 获取物品信息
	var item model.Item
	if err := model.DB.First(&item, itemID).Error; err != nil {
		return nil, err
	}

	// 如果物品可堆叠，尝试找到已有的相同物品
	if item.Stackable {
		var existingItem model.InventoryItem
		result := model.DB.Where("character_id = ? AND item_id = ? AND enhance_level = 0", characterID, itemID).First(&existingItem)
		
		if result.Error == nil {
			// 找到了，增加数量
			newQuantity := existingItem.Quantity + quantity
			if newQuantity > item.MaxStack {
				// 超过最大堆叠数，需要创建新的
				overflow := newQuantity - item.MaxStack
				existingItem.Quantity = item.MaxStack
				model.DB.Save(&existingItem)

				// 创建新的堆叠
				newItem := model.InventoryItem{
					CharacterID: characterID,
					ItemID:      itemID,
					Quantity:    overflow,
				}
				if err := model.DB.Create(&newItem).Error; err != nil {
					return nil, err
				}
				return &newItem, nil
			} else {
				existingItem.Quantity = newQuantity
				if err := model.DB.Save(&existingItem).Error; err != nil {
					return nil, err
				}
				return &existingItem, nil
			}
		}
	}

	// 创建新的仓库项
	inventoryItem := model.InventoryItem{
		CharacterID: characterID,
		ItemID:      itemID,
		Quantity:    quantity,
	}

	if err := model.DB.Create(&inventoryItem).Error; err != nil {
		return nil, err
	}

	return &inventoryItem, nil
}

// RemoveItemFromInventory 从仓库移除物品（内部函数）
func RemoveItemFromInventory(characterID uint, itemID uint, quantity int) error {
	var inventoryItem model.InventoryItem
	result := model.DB.Where("character_id = ? AND item_id = ?", characterID, itemID).First(&inventoryItem)
	
	if result.Error != nil {
		return result.Error
	}

	if inventoryItem.Quantity < quantity {
		return model.DB.Delete(&inventoryItem).Error
	}

	if inventoryItem.Quantity == quantity {
		return model.DB.Delete(&inventoryItem).Error
	}

	inventoryItem.Quantity -= quantity
	return model.DB.Save(&inventoryItem).Error
}

// GetItemQuantity 获取角色拥有的某物品数量（内部函数）
func GetItemQuantity(characterID uint, itemCode string) int {
	var item model.Item
	if err := model.DB.Where("code = ?", itemCode).First(&item).Error; err != nil {
		return 0
	}

	var total int64
	model.DB.Model(&model.InventoryItem{}).
		Where("character_id = ? AND item_id = ?", characterID, item.ID).
		Select("COALESCE(SUM(quantity), 0)").
		Scan(&total)

	return int(total)
}

// HasSufficientItems 检查角色是否有足够的物品（内部函数）
func HasSufficientItems(characterID uint, itemCode string, quantity int) bool {
	return GetItemQuantity(characterID, itemCode) >= quantity
}
