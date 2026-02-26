package main

import (
	"log"

	"plots/internal/config"
	"plots/internal/data"
	"plots/internal/model"
	"plots/internal/router"
)

func main() {
	// 初始化配置
	config.Init()
	log.Println("配置初始化完成")

	// 初始化数据库
	if err := model.InitDB(config.AppConfig.DBPath); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}
	log.Println("数据库初始化完成")

	// 初始化物品数据
	if err := data.InitAllItemData(model.DB); err != nil {
		log.Fatalf("物品数据初始化失败: %v", err)
	}

	// 配置路由
	r := router.SetupRouter()
	log.Println("路由配置完成")

	// 启动服务器
	log.Printf("服务器启动在 http://localhost:%s", config.AppConfig.Port)
	if err := r.Run(":" + config.AppConfig.Port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
