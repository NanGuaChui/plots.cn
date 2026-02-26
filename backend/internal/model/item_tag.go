package model

// ItemTag 物品标签模型
type ItemTag struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Code string `gorm:"size:50;uniqueIndex" json:"code"` // 标签代码：material, crystal, loot, fish, consumable, equipment 等
	Name string `gorm:"size:50" json:"name"`             // 标签名称
}

// TableName 指定表名
func (ItemTag) TableName() string {
	return "item_tags"
}
