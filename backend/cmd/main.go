package main

import (
	"log"

	"poetize/internal/config"
	"poetize/internal/model"
	"poetize/internal/router"
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

	// 配置路由
	r := router.SetupRouter()
	log.Println("路由配置完成")

	// 启动服务器
	log.Printf("服务器启动在 http://localhost:%s", config.AppConfig.Port)
	if err := r.Run(":" + config.AppConfig.Port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
