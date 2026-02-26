package router

import (
	"github.com/gin-gonic/gin"
	"poetize/internal/handler"
	"poetize/internal/middleware"
)

// SetupRouter 配置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS 中间件
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API 路由组
	api := r.Group("/api")
	{
		// 认证相关（无需登录）
		auth := api.Group("/auth")
		{
			auth.POST("/register", handler.Register)
			auth.POST("/login", handler.Login)
		}

		// 需要登录的路由
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// 获取当前用户信息
			protected.GET("/auth/me", handler.GetCurrentUser)

			// 角色相关
			characters := protected.Group("/characters")
			{
				characters.GET("", handler.GetCharacters)
				characters.POST("", handler.CreateCharacter)
				characters.GET("/:id", handler.GetCharacter)
				characters.DELETE("/:id", handler.DeleteCharacter)
				characters.GET("/:id/stats", handler.GetCharacterStats)
			}
		}
	}

	return r
}
