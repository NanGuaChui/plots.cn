package model

import (
	"time"
)

// CharacterType 角色类型
type CharacterType string

const (
	CharacterTypeMain    CharacterType = "main"    // 主号
	CharacterTypeIronman CharacterType = "ironman" // 铁人号
)

// Character 角色模型
type Character struct {
	ID        uint          `gorm:"primaryKey" json:"id"`
	UserID    uint          `gorm:"index" json:"user_id"`     // 所属用户
	Name      string        `gorm:"size:50" json:"name"`      // 角色名
	SlotIndex int           `gorm:"" json:"slot_index"`       // 槽位索引 0-3
	Type      CharacterType `gorm:"size:20" json:"type"`      // 主号/铁人号

	// 基础属性（5项）
	Strength     int `gorm:"default:10" json:"strength"`     // 力量 - 物理攻击、物理技能伤害
	Agility      int `gorm:"default:10" json:"agility"`      // 敏捷 - 速度、闪避率
	Endurance    int `gorm:"default:10" json:"endurance"`    // 耐力 - 生命值、物理防御
	Intelligence int `gorm:"default:10" json:"intelligence"` // 智力 - 魔法攻击、魔法防御、蓝量
	Luck         int `gorm:"default:10" json:"luck"`         // 幸运 - 暴击率、暴击伤害、战利品掉落

	// 战斗等级
	CombatLevel int `gorm:"default:1" json:"combat_level"` // 战斗等级
	CombatExp   int `gorm:"default:0" json:"combat_exp"`   // 战斗经验

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (Character) TableName() string {
	return "characters"
}

// EquipmentBonus 装备加成（占位，后续任务实现）
type EquipmentBonus struct {
	// 基础属性加成
	Strength     int `json:"strength"`
	Agility      int `json:"agility"`
	Endurance    int `json:"endurance"`
	Intelligence int `json:"intelligence"`
	Luck         int `json:"luck"`

	// 直接属性加成
	PhysicalAttack  int     `json:"physical_attack"`
	MagicAttack     int     `json:"magic_attack"`
	PhysicalDefense int     `json:"physical_defense"`
	MagicDefense    int     `json:"magic_defense"`
	MaxHP           int     `json:"max_hp"`
	MaxMP           int     `json:"max_mp"`
	Speed           int     `json:"speed"`
	DodgeRate       float64 `json:"dodge_rate"`
	CritRate        float64 `json:"crit_rate"`
	CritDamage      float64 `json:"crit_damage"`
	LootDropRate    float64 `json:"loot_drop_rate"`
}

// DerivedStats 衍生属性（由基础属性计算得出）
type DerivedStats struct {
	// 攻击属性
	PhysicalAttack int `json:"physical_attack"` // 物理攻击 = 力量 * 2 + 装备加成
	MagicAttack    int `json:"magic_attack"`    // 魔法攻击 = 智力 * 2 + 装备加成

	// 防御属性
	PhysicalDefense int `json:"physical_defense"` // 物理防御 = 耐力 * 1.5 + 装备加成
	MagicDefense    int `json:"magic_defense"`    // 魔法防御 = 智力 * 1 + 装备加成

	// 生命蓝量
	MaxHP int `json:"max_hp"` // 最大生命 = 100 + 耐力 * 10 + 装备加成
	MaxMP int `json:"max_mp"` // 最大蓝量 = 50 + 智力 * 5 + 装备加成

	// 战斗属性
	Speed      int     `json:"speed"`       // 速度 = 敏捷 * 1 + 装备加成
	DodgeRate  float64 `json:"dodge_rate"`  // 闪避率 = 敏捷 * 0.05% + 装备加成
	CritRate   float64 `json:"crit_rate"`   // 暴击率 = 幸运 * 0.1% + 装备加成
	CritDamage float64 `json:"crit_damage"` // 暴击伤害 = 150% + 幸运 * 0.5% + 装备加成

	// 掉落属性
	LootDropRate float64 `json:"loot_drop_rate"` // 战利品掉落率 = 幸运 * 0.1% + 装备加成
}

// CalculateDerivedStats 计算衍生属性
func CalculateDerivedStats(char *Character, equipBonus *EquipmentBonus) DerivedStats {
	// 如果没有装备加成，使用空值
	if equipBonus == nil {
		equipBonus = &EquipmentBonus{}
	}

	// 计算包含装备加成的基础属性
	totalStrength := char.Strength + equipBonus.Strength
	totalAgility := char.Agility + equipBonus.Agility
	totalEndurance := char.Endurance + equipBonus.Endurance
	totalIntelligence := char.Intelligence + equipBonus.Intelligence
	totalLuck := char.Luck + equipBonus.Luck

	return DerivedStats{
		PhysicalAttack:  totalStrength*2 + equipBonus.PhysicalAttack,
		MagicAttack:     totalIntelligence*2 + equipBonus.MagicAttack,
		PhysicalDefense: int(float64(totalEndurance)*1.5) + equipBonus.PhysicalDefense,
		MagicDefense:    totalIntelligence + equipBonus.MagicDefense,
		MaxHP:           100 + totalEndurance*10 + equipBonus.MaxHP,
		MaxMP:           50 + totalIntelligence*5 + equipBonus.MaxMP,
		Speed:           totalAgility + equipBonus.Speed,
		DodgeRate:       float64(totalAgility)*0.0005 + equipBonus.DodgeRate,
		CritRate:        float64(totalLuck)*0.001 + equipBonus.CritRate,
		CritDamage:      1.5 + float64(totalLuck)*0.005 + equipBonus.CritDamage,
		LootDropRate:    float64(totalLuck)*0.001 + equipBonus.LootDropRate,
	}
}

// CharacterWithStats 角色信息（包含衍生属性）
type CharacterWithStats struct {
	Character
	DerivedStats DerivedStats `json:"derived_stats"`
}

// ToCharacterWithStats 转换为包含衍生属性的角色信息
func (c *Character) ToCharacterWithStats(equipBonus *EquipmentBonus) CharacterWithStats {
	return CharacterWithStats{
		Character:    *c,
		DerivedStats: CalculateDerivedStats(c, equipBonus),
	}
}
