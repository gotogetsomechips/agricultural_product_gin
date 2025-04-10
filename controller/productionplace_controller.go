package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"agricultural_product_gin/dto"
	"agricultural_product_gin/model"
	"agricultural_product_gin/service"
)

// ProductionPlaceController 生产地控制器
type ProductionPlaceController struct {
	ProductionPlaceService *service.ProductionPlaceService
}

// NewProductionPlaceController 创建生产地控制器
func NewProductionPlaceController(service *service.ProductionPlaceService) *ProductionPlaceController {
	return &ProductionPlaceController{ProductionPlaceService: service}
}

// Save 新增生产地信息
func (c *ProductionPlaceController) Save(ctx *gin.Context) {
	var place model.ProductionPlace
	if err := ctx.ShouldBindJSON(&place); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 转换模型为DTO
	dto := &dto.ProductionPlaceDTO{
		Address:       place.Address,
		Administrator: place.Administrator,
		Phone:         place.Phone,
	}

	result := c.ProductionPlaceService.CreateProductionPlace(dto)
	ctx.JSON(http.StatusOK, result)
}

// Update 修改生产地信息
func (c *ProductionPlaceController) Update(ctx *gin.Context) {
	var place model.ProductionPlace
	if err := ctx.ShouldBindJSON(&place); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 转换模型为DTO
	dto := &dto.ProductionPlaceDTO{
		ID:            place.ID,
		Address:       place.Address,
		Administrator: place.Administrator,
		Phone:         place.Phone,
	}

	result := c.ProductionPlaceService.UpdateProductionPlace(dto)
	ctx.JSON(http.StatusOK, result)
}

// Delete 删除生产地信息
func (c *ProductionPlaceController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "ID参数错误",
		})
		return
	}

	result := c.ProductionPlaceService.DeleteProductionPlace(id)
	ctx.JSON(http.StatusOK, result)
}

// GetById 根据ID获取生产地信息
func (c *ProductionPlaceController) GetById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "ID参数错误",
		})
		return
	}

	result := c.ProductionPlaceService.GetProductionPlaceByID(id)
	ctx.JSON(http.StatusOK, result)
}

// PageQuery 分页查询生产地信息
func (c *ProductionPlaceController) PageQuery(ctx *gin.Context) {
	var queryDTO dto.ProductionPlacePageQueryDTO
	if err := ctx.ShouldBindJSON(&queryDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	result := c.ProductionPlaceService.PageQueryProductionPlaces(&queryDTO)
	ctx.JSON(http.StatusOK, result)
}

// List 查询所有生产地信息
func (c *ProductionPlaceController) List(ctx *gin.Context) {
	result := c.ProductionPlaceService.GetAllProductionPlaces()
	ctx.JSON(http.StatusOK, result)
}
