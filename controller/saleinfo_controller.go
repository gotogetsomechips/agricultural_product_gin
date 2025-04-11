package controller

import (
	"log"
	"net/http"
	"strconv"

	"agricultural_product_gin/dto"
	"agricultural_product_gin/service"
	"github.com/gin-gonic/gin"
)

// SaleInfoController 销售信息控制器
type SaleInfoController struct {
	service service.SaleInfoService
}

// NewSaleInfoController 创建销售信息控制器
func NewSaleInfoController(service service.SaleInfoService) *SaleInfoController {
	return &SaleInfoController{service: service}
}

// Save 保存销售信息
func (c *SaleInfoController) Save(ctx *gin.Context) {
	var saleInfoDTO dto.SaleInfoDTO
	if err := ctx.ShouldBindJSON(&saleInfoDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	log.Printf("新增销售信息：%+v", saleInfoDTO)
	result := c.service.Save(&saleInfoDTO)
	ctx.JSON(result.Code, result)
}

// Update 更新销售信息
func (c *SaleInfoController) Update(ctx *gin.Context) {
	var saleInfoDTO dto.SaleInfoDTO
	if err := ctx.ShouldBindJSON(&saleInfoDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	log.Printf("修改销售信息：%+v", saleInfoDTO)
	result := c.service.Update(&saleInfoDTO)
	ctx.JSON(result.Code, result)
}

// Delete 删除销售信息
func (c *SaleInfoController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "ID参数错误",
		})
		return
	}

	log.Printf("删除销售信息，ID：%d", id)
	result := c.service.Delete(id)
	ctx.JSON(result.Code, result)
}

// GetByID 根据ID获取销售信息
func (c *SaleInfoController) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "ID参数错误",
		})
		return
	}

	result := c.service.GetByID(id)
	ctx.JSON(result.Code, result)
}

// ListAll 查询所有销售信息
func (c *SaleInfoController) ListAll(ctx *gin.Context) {
	log.Println("查询所有销售信息")
	result := c.service.GetAll()
	ctx.JSON(result.Code, result)
}

// PageQuery 分页查询销售信息
func (c *SaleInfoController) PageQuery(ctx *gin.Context) {
	var queryDTO dto.SaleInfoPageQueryDTO
	if err := ctx.ShouldBindJSON(&queryDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	log.Printf("分页查询销售信息，条件：%+v", queryDTO)
	result := c.service.PageQuery(&queryDTO)
	ctx.JSON(result.Code, result)
}
