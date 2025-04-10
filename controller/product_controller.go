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

// ProductController 产品控制器
type ProductController struct {
	ProductService *service.ProductService
}

// NewProductController 创建产品控制器
func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

// Save 新增产品
func (c *ProductController) Save(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	log.Printf("新增产品：%+v", product)
	result := c.ProductService.CreateProduct(&dto.ProductDTO{
		Name:        product.Name,
		Type:        product.Type,
		Image:       product.Image,
		Description: product.Description,
	})
	ctx.JSON(http.StatusOK, result)
}

// Update 修改产品
func (c *ProductController) Update(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	log.Printf("修改产品：%+v", product)
	result := c.ProductService.UpdateProduct(&dto.ProductDTO{
		ID:          product.ID,
		Name:        product.Name,
		Type:        product.Type,
		Image:       product.Image,
		Description: product.Description,
	})
	ctx.JSON(http.StatusOK, result)
}

// Delete 删除产品
func (c *ProductController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "ID参数错误",
		})
		return
	}

	log.Printf("删除产品，ID：%d", id)
	result := c.ProductService.DeleteProduct(id)
	ctx.JSON(http.StatusOK, result)
}

// GetById 根据ID获取产品
func (c *ProductController) GetById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "ID参数错误",
		})
		return
	}

	result := c.ProductService.GetProductByID(id)
	ctx.JSON(http.StatusOK, result)
}

// PageQuery 分页查询
func (c *ProductController) PageQuery(ctx *gin.Context) {
	var queryDTO dto.ProductPageQueryDTO
	if err := ctx.ShouldBindJSON(&queryDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	log.Printf("分页查询产品，条件：%+v", queryDTO)
	result := c.ProductService.PageQueryProducts(&queryDTO)
	ctx.JSON(http.StatusOK, result)
}

// List 查询所有产品
func (c *ProductController) List(ctx *gin.Context) {
	log.Println("查询所有产品")
	result := c.ProductService.GetAllProducts()
	ctx.JSON(http.StatusOK, result)
}

// GetTypes 获取所有产品类型
func (c *ProductController) GetTypes(ctx *gin.Context) {
	result := c.ProductService.GetProductTypes()
	ctx.JSON(http.StatusOK, result)
}
