package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB(dbPath string) error {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}

	// 自动迁移 - 用户和角色
	err = DB.AutoMigrate(&User{}, &Character{})
	if err != nil {
		return err
	}

	// 自动迁移 - 物品系统
	err = DB.AutoMigrate(&ItemTag{}, &Item{}, &ItemEnhanceConfig{}, &InventoryItem{})
	if err != nil {
		return err
	}

	return nil
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}
