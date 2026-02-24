package main

import (
	"log"
	"os"

	"plots/internal/config"
	"plots/internal/handler"
	"plots/internal/model"
	"plots/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 设置 Gin 模式
	gin.SetMode(cfg.Server.Mode)

	// 初始化数据库
	db, err := repository.InitDB(cfg.Database.Path)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 自动迁移数据库
	if err := model.AutoMigrate(db); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	// 创建仓库实例
	repos := repository.NewRepositories(db)

	// 创建处理器
	h := handler.NewHandler(repos, cfg)

	// 创建路由
	r := gin.Default()

	// 健康检查
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "plots is running",
		})
	})

	// API 路由组
	api := r.Group("/api")
	{
		// 公开接口
		api.GET("/auth/check", h.CheckInit)
		api.POST("/auth/init", h.InitAdmin)
		api.POST("/auth/login", h.Login)
		api.GET("/articles", h.ListArticles)
		api.GET("/articles/:id", h.GetArticle)
		api.GET("/portfolios", h.ListPortfolios)
		api.GET("/portfolios/:id", h.GetPortfolio)

		// 需要认证的接口
		auth := api.Group("/")
		auth.Use(h.AuthMiddleware())
		{
			// 文章管理
			auth.POST("/articles", h.CreateArticle)
			auth.PUT("/articles/:id", h.UpdateArticle)
			auth.DELETE("/articles/:id", h.DeleteArticle)

			// 作品集管理
			auth.POST("/portfolios", h.CreatePortfolio)
			auth.PUT("/portfolios/:id", h.UpdatePortfolio)
			auth.DELETE("/portfolios/:id", h.DeletePortfolio)

			// 用户信息
			auth.GET("/user/profile", h.GetProfile)
			auth.PUT("/user/profile", h.UpdateProfile)
		}
	}

	// 获取端口
	port := cfg.Server.Port
	if port == "" {
		port = os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
	}

	log.Printf("服务启动在端口 %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}
