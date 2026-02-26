package data

import (
	"log"

	"gorm.io/gorm"
	"plots/internal/model"
)

// DefaultTags 默认标签列表
var DefaultTags = []model.ItemTag{
	{Code: "material", Name: "材料"},
	{Code: "crystal", Name: "晶体"},
	{Code: "loot", Name: "战利品"},
	{Code: "fish", Name: "鱼类"},
	{Code: "consumable", Name: "消耗品"},
	{Code: "equipment", Name: "装备"},
	{Code: "weapon", Name: "武器"},
	{Code: "armor", Name: "防具"},
	{Code: "accessory", Name: "饰品"},
	{Code: "tier1", Name: "一级"},
	{Code: "tier2", Name: "二级"},
	{Code: "tier3", Name: "三级"},
	{Code: "wood", Name: "木材"},
	{Code: "ore", Name: "矿石"},
	{Code: "herb", Name: "草药"},
	{Code: "food", Name: "食物"},
}

// ItemDefinition 用于初始化的物品定义
type ItemDefinition struct {
	Item model.Item
	Tags []string // 标签代码列表
}

// DefaultItems 默认物品列表
var DefaultItems = []ItemDefinition{
	// ===== 材料类 - 木材 =====
	{
		Item: model.Item{
			Code:        "oak_log",
			Name:        "橡木原木",
			Description: "从橡树上砍伐获得的木材，是最基础的木材类型。",
			Icon:        "items/oak_log.png",
			Tier:        1,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "wood", "tier1"},
	},
	{
		Item: model.Item{
			Code:        "willow_log",
			Name:        "柳木原木",
			Description: "从柳树上砍伐获得的木材，质地柔韧。",
			Icon:        "items/willow_log.png",
			Tier:        2,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "wood", "tier2"},
	},
	{
		Item: model.Item{
			Code:        "maple_log",
			Name:        "枫木原木",
			Description: "从枫树上砍伐获得的木材，纹理美观。",
			Icon:        "items/maple_log.png",
			Tier:        3,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "wood", "tier3"},
	},

	// ===== 材料类 - 矿石 =====
	{
		Item: model.Item{
			Code:        "copper_ore",
			Name:        "铜矿石",
			Description: "最基础的矿石，可以熔炼成铜锭。",
			Icon:        "items/copper_ore.png",
			Tier:        1,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "ore", "tier1"},
	},
	{
		Item: model.Item{
			Code:        "iron_ore",
			Name:        "铁矿石",
			Description: "常见的矿石，可以熔炼成铁锭。",
			Icon:        "items/iron_ore.png",
			Tier:        2,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "ore", "tier2"},
	},
	{
		Item: model.Item{
			Code:        "gold_ore",
			Name:        "金矿石",
			Description: "珍贵的矿石，可以熔炼成金锭。",
			Icon:        "items/gold_ore.png",
			Tier:        3,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "ore", "tier3"},
	},

	// ===== 材料类 - 金属锭 =====
	{
		Item: model.Item{
			Code:        "copper_bar",
			Name:        "铜锭",
			Description: "由铜矿石熔炼而成的金属锭。",
			Icon:        "items/copper_bar.png",
			Tier:        1,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "tier1"},
	},
	{
		Item: model.Item{
			Code:        "iron_bar",
			Name:        "铁锭",
			Description: "由铁矿石熔炼而成的金属锭。",
			Icon:        "items/iron_bar.png",
			Tier:        2,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "tier2"},
	},

	// ===== 材料类 - 鱼类 =====
	{
		Item: model.Item{
			Code:        "shrimp",
			Name:        "小虾",
			Description: "一种常见的小型甲壳类，可以烹饪食用。",
			Icon:        "items/shrimp.png",
			Tier:        1,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "fish", "tier1"},
	},
	{
		Item: model.Item{
			Code:        "trout",
			Name:        "鳟鱼",
			Description: "一种常见的淡水鱼，肉质鲜美。",
			Icon:        "items/trout.png",
			Tier:        2,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "fish", "tier2"},
	},
	{
		Item: model.Item{
			Code:        "salmon",
			Name:        "三文鱼",
			Description: "一种高级淡水鱼，营养丰富。",
			Icon:        "items/salmon.png",
			Tier:        3,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "fish", "tier3"},
	},

	// ===== 材料类 - 采集品 =====
	{
		Item: model.Item{
			Code:        "herb",
			Name:        "草药",
			Description: "一种常见的草药，可用于制作药剂。",
			Icon:        "items/herb.png",
			Tier:        1,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "herb", "tier1"},
	},
	{
		Item: model.Item{
			Code:        "berry",
			Name:        "浆果",
			Description: "野外采集的浆果，可以直接食用或用于烹饪。",
			Icon:        "items/berry.png",
			Tier:        1,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "food", "tier1"},
	},
	{
		Item: model.Item{
			Code:        "flax",
			Name:        "亚麻",
			Description: "一种纤维植物，可用于制作线。",
			Icon:        "items/flax.png",
			Tier:        2,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "tier2"},
	},

	// ===== 材料类 - 养殖产出 =====
	{
		Item: model.Item{
			Code:        "egg",
			Name:        "鸡蛋",
			Description: "鸡下的蛋，可用于烹饪。",
			Icon:        "items/egg.png",
			Tier:        1,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "food", "tier1"},
	},
	{
		Item: model.Item{
			Code:        "feather",
			Name:        "羽毛",
			Description: "鸡掉落的羽毛，可用于制作箭矢。",
			Icon:        "items/feather.png",
			Tier:        1,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "tier1"},
	},
	{
		Item: model.Item{
			Code:        "wool",
			Name:        "羊毛",
			Description: "从羊身上剪下的羊毛，可用于缝制衣物。",
			Icon:        "items/wool.png",
			Tier:        1,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "tier1"},
	},
	{
		Item: model.Item{
			Code:        "leather",
			Name:        "皮革",
			Description: "从动物身上获得的皮革，可用于缝制防具。",
			Icon:        "items/leather.png",
			Tier:        2,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "tier2"},
	},
	{
		Item: model.Item{
			Code:        "milk",
			Name:        "牛奶",
			Description: "从牛身上挤的牛奶，可用于烹饪。",
			Icon:        "items/milk.png",
			Tier:        2,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "food", "tier2"},
	},

	// ===== 材料类 - 制作中间产物 =====
	{
		Item: model.Item{
			Code:        "thread",
			Name:        "线",
			Description: "由亚麻制成的线，可用于缝制。",
			Icon:        "items/thread.png",
			Tier:        1,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "tier1"},
	},
	{
		Item: model.Item{
			Code:        "cloth",
			Name:        "布料",
			Description: "由线编织而成的布料。",
			Icon:        "items/cloth.png",
			Tier:        2,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "tier2"},
	},
	{
		Item: model.Item{
			Code:        "flour",
			Name:        "面粉",
			Description: "用于烹饪的面粉。",
			Icon:        "items/flour.png",
			Tier:        1,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"material", "food", "tier1"},
	},

	// ===== 消耗品 - 食物 =====
	{
		Item: model.Item{
			Code:        "cooked_shrimp",
			Name:        "烤小虾",
			Description: "烹饪后的小虾，食用可恢复少量生命值。",
			Icon:        "items/cooked_shrimp.png",
			Tier:        1,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"consumable", "food", "tier1"},
	},
	{
		Item: model.Item{
			Code:        "cooked_trout",
			Name:        "烤鳟鱼",
			Description: "烹饪后的鳟鱼，食用可恢复生命值。",
			Icon:        "items/cooked_trout.png",
			Tier:        2,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"consumable", "food", "tier2"},
	},
	{
		Item: model.Item{
			Code:        "berry_pie",
			Name:        "浆果派",
			Description: "用浆果和面粉制作的派，美味可口。",
			Icon:        "items/berry_pie.png",
			Tier:        2,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"consumable", "food", "tier2"},
	},

	// ===== 晶体类 =====
	{
		Item: model.Item{
			Code:        "enhance_crystal_low",
			Name:        "低级强化晶体",
			Description: "用于强化装备的晶体，适用于低级装备。",
			Icon:        "items/enhance_crystal_low.png",
			Tier:        1,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"crystal", "tier1"},
	},
	{
		Item: model.Item{
			Code:        "enhance_crystal_mid",
			Name:        "中级强化晶体",
			Description: "用于强化装备的晶体，适用于中级装备。",
			Icon:        "items/enhance_crystal_mid.png",
			Tier:        2,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"crystal", "tier2"},
	},
	{
		Item: model.Item{
			Code:        "enhance_crystal_high",
			Name:        "高级强化晶体",
			Description: "用于强化装备的晶体，适用于高级装备。",
			Icon:        "items/enhance_crystal_high.png",
			Tier:        3,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"crystal", "tier3"},
	},

	// ===== 保护道具 =====
	{
		Item: model.Item{
			Code:        "protection_scroll_low",
			Name:        "初级保护卷轴",
			Description: "强化失败时只降1级，适用于1-5级强化。",
			Icon:        "items/protection_scroll_low.png",
			Tier:        1,
			Stackable:   true,
			MaxStack:    99,
		},
		Tags: []string{"consumable", "tier1"},
	},
	{
		Item: model.Item{
			Code:        "protection_scroll_mid",
			Name:        "中级保护卷轴",
			Description: "强化失败时只降1级，适用于6-8级强化。",
			Icon:        "items/protection_scroll_mid.png",
			Tier:        2,
			Stackable:   true,
			MaxStack:    99,
		},
		Tags: []string{"consumable", "tier2"},
	},
	{
		Item: model.Item{
			Code:        "protection_scroll_high",
			Name:        "高级保护卷轴",
			Description: "强化失败时只降1级，适用于9-10级强化。",
			Icon:        "items/protection_scroll_high.png",
			Tier:        3,
			Stackable:   true,
			MaxStack:    99,
		},
		Tags: []string{"consumable", "tier3"},
	},

	// ===== 装备类 - 武器 =====
	{
		Item: model.Item{
			Code:            "wooden_sword",
			Name:            "木剑",
			Description:     "一把简陋的木剑，聊胜于无。",
			Icon:            "items/wooden_sword.png",
			Tier:            1,
			Stackable:       false,
			MaxStack:        1,
			EnhanceMaterial: "enhance_crystal_low",
			EnhanceMaxLevel: 5,
		},
		Tags: []string{"equipment", "weapon", "tier1"},
	},
	{
		Item: model.Item{
			Code:            "bronze_sword",
			Name:            "青铜剑",
			Description:     "用青铜锻造的剑，是新手冒险者的标配。",
			Icon:            "items/bronze_sword.png",
			Tier:            1,
			Stackable:       false,
			MaxStack:        1,
			EnhanceMaterial: "enhance_crystal_low",
			EnhanceMaxLevel: 10,
		},
		Tags: []string{"equipment", "weapon", "tier1"},
	},
	{
		Item: model.Item{
			Code:            "iron_sword",
			Name:            "铁剑",
			Description:     "用精铁锻造的剑，锋利而可靠。",
			Icon:            "items/iron_sword.png",
			Tier:            2,
			Stackable:       false,
			MaxStack:        1,
			EnhanceMaterial: "enhance_crystal_mid",
			EnhanceMaxLevel: 10,
		},
		Tags: []string{"equipment", "weapon", "tier2"},
	},

	// ===== 装备类 - 防具 =====
	{
		Item: model.Item{
			Code:            "cloth_armor",
			Name:            "布衣",
			Description:     "一件简单的布制衣物，提供基础防护。",
			Icon:            "items/cloth_armor.png",
			Tier:            1,
			Stackable:       false,
			MaxStack:        1,
			EnhanceMaterial: "enhance_crystal_low",
			EnhanceMaxLevel: 5,
		},
		Tags: []string{"equipment", "armor", "tier1"},
	},
	{
		Item: model.Item{
			Code:            "leather_armor",
			Name:            "皮甲",
			Description:     "用皮革制成的轻型护甲，提供不错的防护。",
			Icon:            "items/leather_armor.png",
			Tier:            2,
			Stackable:       false,
			MaxStack:        1,
			EnhanceMaterial: "enhance_crystal_mid",
			EnhanceMaxLevel: 10,
		},
		Tags: []string{"equipment", "armor", "tier2"},
	},
	{
		Item: model.Item{
			Code:            "wool_hat",
			Name:            "羊毛帽",
			Description:     "用羊毛缝制的帽子，温暖舒适。",
			Icon:            "items/wool_hat.png",
			Tier:            1,
			Stackable:       false,
			MaxStack:        1,
			EnhanceMaterial: "enhance_crystal_low",
			EnhanceMaxLevel: 5,
		},
		Tags: []string{"equipment", "armor", "tier1"},
	},

	// ===== 战利品类 =====
	{
		Item: model.Item{
			Code:        "bone",
			Name:        "骨头",
			Description: "怪物掉落的骨头，可以出售或用于合成。",
			Icon:        "items/bone.png",
			Tier:        1,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"loot", "tier1"},
	},
	{
		Item: model.Item{
			Code:        "slime_gel",
			Name:        "史莱姆凝胶",
			Description: "史莱姆掉落的凝胶，有多种用途。",
			Icon:        "items/slime_gel.png",
			Tier:        1,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"loot", "tier1"},
	},
	{
		Item: model.Item{
			Code:        "goblin_ear",
			Name:        "哥布林耳朵",
			Description: "哥布林掉落的耳朵，证明你击杀了哥布林。",
			Icon:        "items/goblin_ear.png",
			Tier:        1,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"loot", "tier1"},
	},
	{
		Item: model.Item{
			Code:        "wolf_pelt",
			Name:        "狼皮",
			Description: "野狼掉落的皮毛，可以出售或用于制作。",
			Icon:        "items/wolf_pelt.png",
			Tier:        2,
			Stackable:   true,
			MaxStack:    999,
		},
		Tags: []string{"loot", "tier2"},
	},
}

// InitItemTags 初始化物品标签
func InitItemTags(db *gorm.DB) error {
	for _, tag := range DefaultTags {
		var existing model.ItemTag
		result := db.Where("code = ?", tag.Code).First(&existing)
		if result.Error == gorm.ErrRecordNotFound {
			if err := db.Create(&tag).Error; err != nil {
				return err
			}
			log.Printf("创建标签: %s (%s)", tag.Name, tag.Code)
		}
	}
	return nil
}

// InitItems 初始化物品
func InitItems(db *gorm.DB) error {
	// 先获取所有标签的映射
	var tags []model.ItemTag
	if err := db.Find(&tags).Error; err != nil {
		return err
	}
	tagMap := make(map[string]model.ItemTag)
	for _, tag := range tags {
		tagMap[tag.Code] = tag
	}

	// 创建物品
	for _, itemDef := range DefaultItems {
		var existing model.Item
		result := db.Where("code = ?", itemDef.Item.Code).First(&existing)
		if result.Error == gorm.ErrRecordNotFound {
			// 关联标签
			var itemTags []model.ItemTag
			for _, tagCode := range itemDef.Tags {
				if tag, ok := tagMap[tagCode]; ok {
					itemTags = append(itemTags, tag)
				}
			}
			itemDef.Item.Tags = itemTags

			if err := db.Create(&itemDef.Item).Error; err != nil {
				return err
			}
			log.Printf("创建物品: %s (%s)", itemDef.Item.Name, itemDef.Item.Code)
		}
	}
	return nil
}

// InitAllItemData 初始化所有物品相关数据
func InitAllItemData(db *gorm.DB) error {
	log.Println("开始初始化物品数据...")

	if err := InitItemTags(db); err != nil {
		return err
	}

	if err := InitItems(db); err != nil {
		return err
	}

	log.Println("物品数据初始化完成")
	return nil
}
