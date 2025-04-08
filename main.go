package main

import (
	"log"
	"time"

	"agricultural_product_gin/config"
	"agricultural_product_gin/controller"
	"agricultural_product_gin/middleware"
	"agricultural_product_gin/repository"
	"agricultural_product_gin/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	db := config.GetDB()
	defer db.Close()

	// 创建依赖
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// 创建Gin引擎
	r := gin.Default()
	
	// 配置CORS中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 用户相关路由
	userGroup := r.Group("/user")
	{
		// 公开API
		userGroup.POST("/register", userController.Register)
		userGroup.POST("/login", userController.Login)

		// 需要认证的API
		authGroup := userGroup.Group("/")
		authGroup.Use(middleware.JWTMiddleware())
		{
			authGroup.POST("/logout", userController.Logout)
			authGroup.GET("/userInfo", userController.GetUserInfo)
			authGroup.PUT("/update", userController.Update)
			authGroup.PUT("/editPassword", userController.EditPassword)
		}
	}

	// 启动服务器
	log.Println("服务器启动在 :8080 端口")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}
