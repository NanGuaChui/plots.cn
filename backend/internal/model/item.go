package model

// Item 物品定义模型
type Item struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Code        string `gorm:"size:50;uniqueIndex" json:"code"` // 物品代码，如 "oak_log", "iron_ore"
	Name        string `gorm:"size:100" json:"name"`            // 物品名称
	Description string `gorm:"size:500" json:"description"`     // 物品描述
	Icon        string `gorm:"size:200" json:"icon"`            // 图标路径
	Tier        int    `gorm:"default:1" json:"tier"`           // 物品等级

	// 堆叠属性
	Stackable bool `gorm:"default:true" json:"stackable"` // 是否可堆叠
	MaxStack  int  `gorm:"default:999" json:"max_stack"`  // 最大堆叠数

	// 强化相关（装备类物品）
	EnhanceMaterial string `gorm:"size:50" json:"enhance_material"`   // 强化所需材料代码
	EnhanceMaxLevel int    `gorm:"default:0" json:"enhance_max_level"` // 强化上限

	// 标签关联（多对多）
	Tags []ItemTag `gorm:"many2many:item_item_tags" json:"tags"`
}

// TableName 指定表名
func (Item) TableName() string {
	return "items"
}

// ItemEnhanceConfig 物品强化配置
type ItemEnhanceConfig struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	ItemID       uint    `gorm:"index" json:"item_id"`       // 物品ID
	Level        int     `gorm:"" json:"level"`              // 强化等级
	FailureRate  float64 `gorm:"" json:"failure_rate"`       // 失败率 0.0-1.0
	MaterialCost int     `gorm:"" json:"material_cost"`      // 材料消耗数量
}

// TableName 指定表名
func (ItemEnhanceConfig) TableName() string {
	return "item_enhance_configs"
}
