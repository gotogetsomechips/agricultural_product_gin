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

	// 创建用户相关依赖
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// 创建产品相关依赖
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	// 创建文件上传控制器
	uploadController := controller.NewUploadController()

	// 创建Gin引擎
	r := gin.Default()

	// 配置CORS中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3030", "http://75249acd.r3.cpolar.top"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 注册静态文件路由 - 用于访问上传的图片
	uploadController.RegisterStaticRoutes(r)

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

	productGroup := r.Group("/product")
	{
		// 路由映射
		productGroup.POST("", productController.Save)           // 新增
		productGroup.DELETE("/:id", productController.Delete)   // 删除
		productGroup.GET("/:id", productController.GetById)     // 根据id查询
		productGroup.POST("/page", productController.PageQuery) // 分页查询
		productGroup.PUT("", productController.Update)          // 修改
		productGroup.GET("/list", productController.List)       // 查询所有
		productGroup.GET("/types", productController.GetTypes)  // 获取所有产品类型
	}

	// 创建生产信息相关依赖
	productionRepo := repository.NewProductionRepository(db)
	productionService := service.NewProductionService(productionRepo)
	productionController := controller.NewProductionController(productionService)

	// 生产信息路由组
	productionGroup := r.Group("/productinfo")
	{
		productionGroup.POST("", productionController.Save)           // 新增
		productionGroup.DELETE("/:id", productionController.Delete)   // 删除
		productionGroup.GET("/:id", productionController.GetById)     // 根据id查询
		productionGroup.POST("/page", productionController.PageQuery) // 分页查询
		productionGroup.PUT("", productionController.Update)          // 修改
		productionGroup.GET("/list", productionController.List)       // 查询所有
	}
	// 在main.go中添加以下代码

	// 创建生产地相关依赖
	productionPlaceRepo := repository.NewProductionPlaceRepository(db)
	productionPlaceService := service.NewProductionPlaceService(productionPlaceRepo)
	productionPlaceController := controller.NewProductionPlaceController(productionPlaceService)

	// 生产地路由组
	productionPlaceGroup := r.Group("/productplace")
	{
		productionPlaceGroup.POST("", productionPlaceController.Save)           // 新增
		productionPlaceGroup.DELETE("/:id", productionPlaceController.Delete)   // 删除
		productionPlaceGroup.GET("/:id", productionPlaceController.GetById)     // 根据id查询
		productionPlaceGroup.POST("/page", productionPlaceController.PageQuery) // 分页查询
		productionPlaceGroup.PUT("", productionPlaceController.Update)          // 修改
		productionPlaceGroup.GET("/list", productionPlaceController.List)       // 查询所有
	}

	// 创建公司相关依赖
	companyRepo := repository.NewCompanyRepository(db)
	companyService := service.NewCompanyService(companyRepo)
	companyController := controller.NewCompanyController(companyService)

	// 公司路由组
	companyGroup := r.Group("/company")
	{
		companyGroup.POST("", companyController.Save)           // 新增
		companyGroup.DELETE("/:id", companyController.Delete)   // 删除
		companyGroup.GET("/:id", companyController.GetByID)     // 根据id查询
		companyGroup.POST("/page", companyController.PageQuery) // 分页查询
		companyGroup.PUT("", companyController.Update)          // 修改
		companyGroup.GET("/list", companyController.ListAll)    // 查询所有
	}

	// 创建物流相关依赖
	logisticsRepo := repository.NewLogisticsRepository(db)
	logisticsService := service.NewLogisticsService(logisticsRepo)
	logisticsController := controller.NewLogisticsController(logisticsService)

	// 物流路由组
	logisticsGroup := r.Group("/logistics")
	{
		logisticsGroup.POST("", logisticsController.Save)                      // 新增
		logisticsGroup.DELETE("/:id", logisticsController.Delete)              // 删除
		logisticsGroup.GET("/:id", logisticsController.GetById)                // 根据id查询
		logisticsGroup.POST("/page", logisticsController.PageQuery)            // 分页查询
		logisticsGroup.PUT("", logisticsController.Update)                     // 修改
		logisticsGroup.GET("/list", logisticsController.List)                  // 查询所有
		logisticsGroup.PUT("/confirm/:id", logisticsController.ConfirmReceipt) // 确认收货
	}

	// 创建销售地相关依赖
	salePlaceRepo := repository.NewSalePlaceRepository(db)
	salePlaceService := service.NewSalePlaceService(salePlaceRepo)
	salePlaceController := controller.NewSalePlaceController(salePlaceService)

	// 销售地路由组
	salePlaceGroup := r.Group("/saleplace")
	{
		salePlaceGroup.POST("", salePlaceController.Save)           // 新增
		salePlaceGroup.DELETE("/:id", salePlaceController.Delete)   // 删除
		salePlaceGroup.GET("/:id", salePlaceController.GetByID)     // 根据id查询
		salePlaceGroup.POST("/page", salePlaceController.PageQuery) // 分页查询
		salePlaceGroup.PUT("", salePlaceController.Update)          // 修改
		salePlaceGroup.GET("/list", salePlaceController.ListAll)    // 查询所有
	}
	// 创建销售信息相关依赖
	saleInfoRepo := repository.NewSaleInfoRepository(db)
	saleInfoService := service.NewSaleInfoService(saleInfoRepo)
	saleInfoController := controller.NewSaleInfoController(saleInfoService)

	// 销售信息路由组
	saleInfoGroup := r.Group("/saleinfo")
	{
		saleInfoGroup.POST("", saleInfoController.Save)           // 新增
		saleInfoGroup.PUT("", saleInfoController.Update)          // 修改
		saleInfoGroup.DELETE("/:id", saleInfoController.Delete)   // 删除
		saleInfoGroup.GET("/:id", saleInfoController.GetByID)     // 根据id查询
		saleInfoGroup.GET("/list", saleInfoController.ListAll)    // 查询所有
		saleInfoGroup.POST("/page", saleInfoController.PageQuery) // 分页查询
	}
	r.POST("/upload", uploadController.Upload)

	// 启动服务器
	log.Println("服务器启动在 :8080 端口")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}
