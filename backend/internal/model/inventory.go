package model

import "time"

// InventoryItem 仓库物品模型
type InventoryItem struct {
	ID           uint `gorm:"primaryKey" json:"id"`
	CharacterID  uint `gorm:"index" json:"character_id"`   // 所属角色
	ItemID       uint `gorm:"index" json:"item_id"`        // 物品ID
	Quantity     int  `gorm:"default:1" json:"quantity"`   // 数量
	EnhanceLevel int  `gorm:"default:0" json:"enhance_level"` // 强化等级（装备类）

	// 关联
	Item Item `gorm:"foreignKey:ItemID" json:"item"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (InventoryItem) TableName() string {
	return "inventory_items"
}

// InventoryItemWithDetails 仓库物品详情（包含物品信息）
type InventoryItemWithDetails struct {
	InventoryItem
	ItemName string    `json:"item_name"`
	ItemCode string    `json:"item_code"`
	ItemIcon string    `json:"item_icon"`
	Tags     []ItemTag `json:"tags"`
}
