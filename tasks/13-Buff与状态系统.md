# 任务 13：Buff 与状态系统

## 📋 任务目标
实现 Buff/Debuff 系统，包括状态定义、叠加规则、效果触发。

## 📁 需要创建/修改的文件

### 后端文件
```
backend/internal/model/
└── buff.go             # Buff 定义模型

backend/internal/battle/
├── buff_manager.go     # Buff 管理器
└── buff_effect.go      # Buff 效果处理

backend/internal/data/
└── buffs_init.go       # Buff 初始化数据
```

## 📝 详细实现要求

### 1. Buff 定义模型
```go
type BuffType string
type StackType string

const (
    BuffTypePositive BuffType = "positive"  // 增益
    BuffTypeNegative BuffType = "negative"  // 减益
)

const (
    StackRefresh StackType = "refresh"  // 刷新持续时间
    StackAdd     StackType = "add"      // 叠加层数
    StackNone    StackType = "none"     // 不可叠加
)

type Buff struct {
    ID          uint      `gorm:"primaryKey"`
    Code        string    `gorm:"size:50;uniqueIndex"`
    Name        string    `gorm:"size:100"`
    Description string    `gorm:"size:500"`
    Type        BuffType  `gorm:"size:20"`
    StackType   StackType `gorm:"size:20"`
    MaxStacks   int       `gorm:"default:1"`     // 最大叠加层数
    
    // 效果配置（JSON）
    EffectConfig string   `gorm:"type:text"`
    
    Icon        string    `gorm:"size:200"`
}
```

### 2. Buff 实例模型 (运行时)
```go
type BuffInstance struct {
    BuffID      uint   `json:"buff_id"`
    BuffCode    string `json:"buff_code"`
    Name        string `json:"name"`
    SourceID    string `json:"source_id"`    // 施加者ID
    Stacks      int    `json:"stacks"`       // 当前层数
    Duration    int    `json:"duration"`     // 剩余回合数
    EffectValue int    `json:"effect_value"` // 效果数值
}
```

### 3. Buff 效果类型
```go
type BuffEffectType string

const (
    // 属性修改
    EffectModAttack    BuffEffectType = "mod_attack"     // 修改攻击
    EffectModDefense   BuffEffectType = "mod_defense"    // 修改防御
    EffectModSpeed     BuffEffectType = "mod_speed"      // 修改速度
    EffectModCritRate  BuffEffectType = "mod_crit_rate"  // 修改暴击率
    
    // 持续效果
    EffectDOT          BuffEffectType = "dot"            // 持续伤害
    EffectHOT          BuffEffectType = "hot"            // 持续治疗
    
    // 控制效果
    EffectStun         BuffEffectType = "stun"           // 眩晕
    EffectFreeze       BuffEffectType = "freeze"         // 冰冻
    EffectSilence      BuffEffectType = "silence"        // 沉默
    
    // 特殊效果
    EffectReflect      BuffEffectType = "reflect"        // 反伤
    EffectShield       BuffEffectType = "shield"         // 护盾
    EffectImmune       BuffEffectType = "immune"         // 免疫
)

type BuffEffect struct {
    Type       BuffEffectType `json:"type"`
    Value      float64        `json:"value"`       // 效果数值
    IsPercent  bool           `json:"is_percent"`  // 是否百分比
}
```

### 4. Buff 管理器
```go
type BuffManager struct {
    state *BattleState
}

// ApplyBuff 施加 Buff
func (m *BuffManager) ApplyBuff(targetID string, buff Buff, sourceID string, duration int, value int) {
    target := m.getUnit(targetID)
    
    // 检查是否已存在
    existing := m.findExistingBuff(target, buff.Code)
    
    if existing != nil {
        // 根据叠加规则处理
        switch buff.StackType {
        case StackRefresh:
            existing.Duration = duration
        case StackAdd:
            if existing.Stacks < buff.MaxStacks {
                existing.Stacks++
            }
            existing.Duration = duration
        case StackNone:
            // 不做处理
            return
        }
    } else {
        // 添加新 Buff
        instance := BuffInstance{
            BuffID:      buff.ID,
            BuffCode:    buff.Code,
            Name:        buff.Name,
            SourceID:    sourceID,
            Stacks:      1,
            Duration:    duration,
            EffectValue: value,
        }
        
        if buff.Type == BuffTypePositive {
            target.Buffs = append(target.Buffs, instance)
        } else {
            target.Debuffs = append(target.Debuffs, instance)
        }
    }
}

// RemoveBuff 移除 Buff
func (m *BuffManager) RemoveBuff(targetID string, buffCode string)

// ProcessBuffs 处理回合开始/结束的 Buff 效果
func (m *BuffManager) ProcessBuffs(phase string) []BattleEvent {
    events := []BattleEvent{}
    
    for _, unit := range m.getAllUnits() {
        // 处理持续伤害
        for _, debuff := range unit.Debuffs {
            effect := m.getBuffEffect(debuff.BuffCode)
            if effect.Type == EffectDOT {
                damage := m.calculateDOTDamage(debuff)
                unit.CurrentHP -= damage
                events = append(events, BattleEvent{
                    Type:   "dot_damage",
                    Target: unit.ID,
                    Value:  damage,
                    Buff:   debuff.BuffCode,
                })
            }
        }
        
        // 处理持续治疗
        for _, buff := range unit.Buffs {
            effect := m.getBuffEffect(buff.BuffCode)
            if effect.Type == EffectHOT {
                heal := m.calculateHOTHeal(buff)
                unit.CurrentHP = min(unit.CurrentHP + heal, unit.MaxHP)
                events = append(events, BattleEvent{
                    Type:   "hot_heal",
                    Target: unit.ID,
                    Value:  heal,
                    Buff:   buff.BuffCode,
                })
            }
        }
        
        // 减少持续时间
        m.tickBuffDurations(unit)
    }
    
    return events
}
```

### 5. 属性修改计算
```go
// GetModifiedStats 获取 Buff 修改后的属性
func (m *BuffManager) GetModifiedStats(unit *BattleUnit) UnitStats {
    stats := unit.BaseStats()
    
    // 应用所有 Buff 效果
    for _, buff := range unit.Buffs {
        effect := m.getBuffEffect(buff.BuffCode)
        m.applyStatModifier(&stats, effect, buff.Stacks)
    }
    
    for _, debuff := range unit.Debuffs {
        effect := m.getBuffEffect(debuff.BuffCode)
        m.applyStatModifier(&stats, effect, debuff.Stacks)
    }
    
    return stats
}

func (m *BuffManager) applyStatModifier(stats *UnitStats, effect BuffEffect, stacks int) {
    value := effect.Value * float64(stacks)
    
    switch effect.Type {
    case EffectModAttack:
        if effect.IsPercent {
            stats.Attack = int(float64(stats.Attack) * (1 + value/100))
        } else {
            stats.Attack += int(value)
        }
    // ... 其他属性
    }
}
```

### 6. 控制效果检查
```go
// CanAct 检查单位是否可以行动
func (m *BuffManager) CanAct(unit *BattleUnit) bool {
    for _, debuff := range unit.Debuffs {
        effect := m.getBuffEffect(debuff.BuffCode)
        if effect.Type == EffectStun || effect.Type == EffectFreeze {
            return false
        }
    }
    return true
}

// CanUseSkill 检查是否可以使用技能
func (m *BuffManager) CanUseSkill(unit *BattleUnit) bool {
    for _, debuff := range unit.Debuffs {
        effect := m.getBuffEffect(debuff.BuffCode)
        if effect.Type == EffectSilence {
            return false
        }
    }
    return true
}
```

### 7. 初始 Buff 数据
```go
var DefaultBuffs = []Buff{
    // 减益
    {Code: "burn", Name: "灼烧", Type: BuffTypeNegative, StackType: StackAdd, MaxStacks: 5,
     EffectConfig: `{"type": "dot", "value": 10, "is_percent": false}`},
    {Code: "poison", Name: "中毒", Type: BuffTypeNegative, StackType: StackRefresh, MaxStacks: 1,
     EffectConfig: `{"type": "dot", "value": 5, "is_percent": true}`},
    {Code: "freeze", Name: "冰冻", Type: BuffTypeNegative, StackType: StackNone, MaxStacks: 1,
     EffectConfig: `{"type": "freeze"}`},
    {Code: "stun", Name: "眩晕", Type: BuffTypeNegative, StackType: StackNone, MaxStacks: 1,
     EffectConfig: `{"type": "stun"}`},
    {Code: "weakness", Name: "虚弱", Type: BuffTypeNegative, StackType: StackRefresh, MaxStacks: 1,
     EffectConfig: `{"type": "mod_attack", "value": -20, "is_percent": true}`},
    
    // 增益
    {Code: "regen", Name: "再生", Type: BuffTypePositive, StackType: StackRefresh, MaxStacks: 1,
     EffectConfig: `{"type": "hot", "value": 5, "is_percent": true}`},
    {Code: "attack_up", Name: "攻击提升", Type: BuffTypePositive, StackType: StackAdd, MaxStacks: 3,
     EffectConfig: `{"type": "mod_attack", "value": 10, "is_percent": true}`},
    {Code: "defense_up", Name: "防御提升", Type: BuffTypePositive, StackType: StackRefresh, MaxStacks: 1,
     EffectConfig: `{"type": "mod_defense", "value": 20, "is_percent": true}`},
    {Code: "shield", Name: "护盾", Type: BuffTypePositive, StackType: StackNone, MaxStacks: 1,
     EffectConfig: `{"type": "shield", "value": 100}`},
}
```

## ✅ 验证检查点

完成后请验证：
1. [ ] Buff 正确施加到目标
2. [ ] 刷新型 Buff 正确刷新时间
3. [ ] 叠加型 Buff 正确增加层数
4. [ ] 持续伤害/治疗正确触发
5. [ ] 属性修改正确计算
6. [ ] 控制效果正确阻止行动
7. [ ] Buff 到期后正确移除
8. [ ] 前端正确显示 Buff 图标和层数

## 🔗 依赖关系
- 依赖：任务 11（战斗回合机制）
- 依赖：任务 12（技能系统，技能附带 Buff）
