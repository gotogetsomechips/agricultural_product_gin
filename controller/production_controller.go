package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"agricultural_product_gin/dto"
	"agricultural_product_gin/model"
	"agricultural_product_gin/service"
)

// ProductionController 生产信息控制器
type ProductionController struct {
	ProductionService *service.ProductionService
}

// NewProductionController 创建生产信息控制器
func NewProductionController(service *service.ProductionService) *ProductionController {
	return &ProductionController{ProductionService: service}
}

// Save 新增生产信息
func (c *ProductionController) Save(ctx *gin.Context) {
	var production model.ProductionInfo
	if err := ctx.ShouldBindJSON(&production); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 转换模型为DTO
	dto := &dto.ProductionDTO{
		ProductID:      production.ProductID,
		ProductPlaceID: production.ProductPlaceID,
		SeedSource:     production.SeedSource,
		Description:    production.Description,
		PlantingDate:   production.PlantingDate,
		HarvestDate:    production.HarvestDate,
	}

	result := c.ProductionService.CreateProduction(dto)
	ctx.JSON(http.StatusOK, result)
}

// Update 修改生产信息
func (c *ProductionController) Update(ctx *gin.Context) {
	var production model.ProductionInfo
	if err := ctx.ShouldBindJSON(&production); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 转换模型为DTO
	dto := &dto.ProductionDTO{
		ID:             production.ID,
		ProductID:      production.ProductID,
		ProductPlaceID: production.ProductPlaceID,
		SeedSource:     production.SeedSource,
		Description:    production.Description,
		PlantingDate:   production.PlantingDate,
		HarvestDate:    production.HarvestDate,
	}

	result := c.ProductionService.UpdateProduction(dto)
	ctx.JSON(http.StatusOK, result)
}

// Delete 删除生产信息
func (c *ProductionController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "ID参数错误",
		})
		return
	}

	result := c.ProductionService.DeleteProduction(id)
	ctx.JSON(http.StatusOK, result)
}

// GetById 根据ID获取生产信息
func (c *ProductionController) GetById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "ID参数错误",
		})
		return
	}

	result := c.ProductionService.GetProductionByID(id)
	ctx.JSON(http.StatusOK, result)
}

// PageQuery 分页查询生产信息
func (c *ProductionController) PageQuery(ctx *gin.Context) {
	var queryDTO dto.ProductionPageQueryDTO
	if err := ctx.ShouldBindJSON(&queryDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	result := c.ProductionService.PageQueryProductions(&queryDTO)
	ctx.JSON(http.StatusOK, result)
}

// List 查询所有生产信息
func (c *ProductionController) List(ctx *gin.Context) {
	result := c.ProductionService.GetAllProductions()
	ctx.JSON(http.StatusOK, result)
}
