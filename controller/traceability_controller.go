package controller

import (
	"net/http"
	"strconv"

	"agricultural_product_gin/service"
	"github.com/gin-gonic/gin"
)

// TraceabilityController 处理所有溯源相关的端点
type TraceabilityController struct {
	productionService *service.ProductionService
	productService    *service.ProductService
	logisticsService  *service.LogisticsService
	saleInfoService   service.SaleInfoService // 注意这里是接口类型，不是指针
}

// NewTraceabilityController 创建一个新的溯源控制器实例
func NewTraceabilityController(
	productionService *service.ProductionService,
	productService *service.ProductService,
	logisticsService *service.LogisticsService,
	saleInfoService service.SaleInfoService,
) *TraceabilityController {
	return &TraceabilityController{
		productionService: productionService,
		productService:    productService,
		logisticsService:  logisticsService,
		saleInfoService:   saleInfoService,
	}
}

// RegisterRoutes 注册所有溯源相关路由
func (tc *TraceabilityController) RegisterRoutes(router *gin.RouterGroup) {
	traceabilityGroup := router.Group("/traceability")
	{
		traceabilityGroup.GET("/productinfo/:id", tc.GetProductInfo)
		traceabilityGroup.GET("/saleinfo/:id", tc.GetSaleInfo)
		traceabilityGroup.GET("/logistics/:id", tc.GetLogistics)
		traceabilityGroup.GET("/product/:id", tc.GetProduct)
	}
}

// GetProductInfo 通过ID获取生产信息
// @Summary 查询生产信息
// @Router /traceability/productinfo/{id} [get]
func (tc *TraceabilityController) GetProductInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "ID格式无效", "data": nil})
		return
	}

	result := tc.productionService.GetProductionByID(id)
	c.JSON(http.StatusOK, result)
}

// GetSaleInfo 通过ID获取销售信息
// @Summary 查询销售信息
// @Router /traceability/saleinfo/{id} [get]
func (tc *TraceabilityController) GetSaleInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "ID格式无效", "data": nil})
		return
	}

	result := tc.saleInfoService.GetByID(id)
	c.JSON(http.StatusOK, result)
}

// GetLogistics 通过ID获取物流信息
// @Summary 查询物流信息
// @Router /traceability/logistics/{id} [get]
func (tc *TraceabilityController) GetLogistics(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "ID格式无效", "data": nil})
		return
	}

	logistics, err := tc.logisticsService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": err.Error(), "data": nil})
		return
	}

	if logistics == nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "msg": "物流信息不存在", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": logistics})
}

// GetProduct 通过ID获取产品信息
// @Summary 查询产品信息
// @Router /traceability/product/{id} [get]
func (tc *TraceabilityController) GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "ID格式无效", "data": nil})
		return
	}

	result := tc.productService.GetProductByID(id)
	c.JSON(http.StatusOK, result)
}
