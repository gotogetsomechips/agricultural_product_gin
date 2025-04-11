package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"agricultural_product_gin/dto"
	"agricultural_product_gin/model"
	"agricultural_product_gin/service"
)

// SalePlaceController 销售地控制器
type SalePlaceController struct {
	SalePlaceService *service.SalePlaceService
}

// NewSalePlaceController 创建销售地控制器
func NewSalePlaceController(salePlaceService *service.SalePlaceService) *SalePlaceController {
	return &SalePlaceController{SalePlaceService: salePlaceService}
}

// Save 新增销售地
func (c *SalePlaceController) Save(ctx *gin.Context) {
	var salePlace model.SalePlace
	if err := ctx.ShouldBindJSON(&salePlace); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	log.Printf("新增销售地：%+v", salePlace)
	result := c.SalePlaceService.CreateSalePlace(&dto.SalePlaceDTO{
		ID:            salePlace.ID,
		Address:       salePlace.Address,
		Administrator: salePlace.Administrator,
		Phone:         salePlace.Phone,
	})
	ctx.JSON(http.StatusOK, result)
}

// Update 修改销售地
func (c *SalePlaceController) Update(ctx *gin.Context) {
	var salePlace model.SalePlace
	if err := ctx.ShouldBindJSON(&salePlace); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	log.Printf("修改销售地：%+v", salePlace)
	result := c.SalePlaceService.UpdateSalePlace(&dto.SalePlaceDTO{
		ID:            salePlace.ID,
		Address:       salePlace.Address,
		Administrator: salePlace.Administrator,
		Phone:         salePlace.Phone,
	})
	ctx.JSON(http.StatusOK, result)
}

// Delete 删除销售地
func (c *SalePlaceController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "ID参数错误",
		})
		return
	}

	log.Printf("删除销售地，ID：%d", id)
	result := c.SalePlaceService.DeleteSalePlace(id)
	ctx.JSON(http.StatusOK, result)
}

// GetByID 根据ID获取销售地
func (c *SalePlaceController) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "ID参数错误",
		})
		return
	}

	result := c.SalePlaceService.GetSalePlaceByID(id)
	ctx.JSON(http.StatusOK, result)
}

// PageQuery 分页查询
func (c *SalePlaceController) PageQuery(ctx *gin.Context) {
	var queryDTO dto.SalePlacePageQueryDTO
	if err := ctx.ShouldBindJSON(&queryDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	log.Printf("分页查询销售地，条件：%+v", queryDTO)
	result := c.SalePlaceService.PageQuerySalePlaces(&queryDTO)
	ctx.JSON(http.StatusOK, result)
}

// ListAll 查询所有销售地
func (c *SalePlaceController) ListAll(ctx *gin.Context) {
	log.Println("查询所有销售地")
	result := c.SalePlaceService.GetAllSalePlaces()
	ctx.JSON(http.StatusOK, result)
}
