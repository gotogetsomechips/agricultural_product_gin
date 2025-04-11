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

// CompanyController 公司控制器
type CompanyController struct {
	CompanyService *service.CompanyService
}

// NewCompanyController 创建公司控制器
func NewCompanyController(companyService *service.CompanyService) *CompanyController {
	return &CompanyController{CompanyService: companyService}
}

// Save 新增公司
func (c *CompanyController) Save(ctx *gin.Context) {
	var company model.Company
	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	log.Printf("新增公司：%+v", company)
	result := c.CompanyService.CreateCompany(&dto.CompanyDTO{
		Name:          company.Name,
		Address:       company.Address,
		Administrator: company.Administrator,
		Phone:         company.Phone,
	})
	ctx.JSON(http.StatusOK, result)
}

// Update 修改公司
func (c *CompanyController) Update(ctx *gin.Context) {
	var company model.Company
	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	log.Printf("修改公司：%+v", company)
	result := c.CompanyService.UpdateCompany(&dto.CompanyDTO{
		ID:            company.ID,
		Name:          company.Name,
		Address:       company.Address,
		Administrator: company.Administrator,
		Phone:         company.Phone,
	})
	ctx.JSON(http.StatusOK, result)
}

// Delete 删除公司
func (c *CompanyController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "ID参数错误",
		})
		return
	}

	log.Printf("删除公司，ID：%d", id)
	result := c.CompanyService.DeleteCompany(id)
	ctx.JSON(http.StatusOK, result)
}

// GetByID 根据ID获取公司
func (c *CompanyController) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "ID参数错误",
		})
		return
	}

	result := c.CompanyService.GetCompanyByID(id)
	ctx.JSON(http.StatusOK, result)
}

// PageQuery 分页查询
func (c *CompanyController) PageQuery(ctx *gin.Context) {
	var queryDTO dto.CompanyPageQueryDTO
	if err := ctx.ShouldBindJSON(&queryDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	log.Printf("分页查询公司，条件：%+v", queryDTO)
	result := c.CompanyService.PageQueryCompanies(&queryDTO)
	ctx.JSON(http.StatusOK, result)
}

// ListAll 查询所有公司
func (c *CompanyController) ListAll(ctx *gin.Context) {
	log.Println("查询所有公司")
	result := c.CompanyService.GetAllCompanies()
	ctx.JSON(http.StatusOK, result)
}
