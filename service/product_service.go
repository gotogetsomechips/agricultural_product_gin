package service

import (
	"log"

	"agricultural_product_gin/dto"
	"agricultural_product_gin/model"
	"agricultural_product_gin/repository"
)

// ProductService 产品服务
type ProductService struct {
	ProductRepo *repository.ProductRepository
}

// NewProductService 创建产品服务
func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: productRepo}
}

// CreateProduct 创建产品
func (s *ProductService) CreateProduct(productDTO *dto.ProductDTO) *dto.Result {
	// 转换DTO为模型
	product := &model.Product{
		Name:        productDTO.Name,
		Type:        productDTO.Type,
		Image:       productDTO.Image,
		Description: productDTO.Description,
	}

	// 保存产品
	id, err := s.ProductRepo.Save(product)
	if err != nil {
		log.Println("创建产品失败:", err)
		return errorResult(500, "创建产品失败")
	}

	// 返回创建成功的产品ID
	return successResult("添加成功", id)
}

// UpdateProduct 更新产品
func (s *ProductService) UpdateProduct(productDTO *dto.ProductDTO) *dto.Result {
	// 检查产品是否存在
	existingProduct, err := s.ProductRepo.GetByID(productDTO.ID)
	if err != nil {
		log.Println("查询产品失败:", err)
		return errorResult(500, "系统错误")
	}

	if existingProduct == nil {
		return errorResult(404, "产品不存在")
	}

	// 转换DTO为模型
	product := &model.Product{
		ID:          productDTO.ID,
		Name:        productDTO.Name,
		Type:        productDTO.Type,
		Image:       productDTO.Image,
		Description: productDTO.Description,
	}

	// 更新产品
	err = s.ProductRepo.Update(product)
	if err != nil {
		log.Println("更新产品失败:", err)
		return errorResult(500, "更新失败")
	}

	return successResult("更新成功", nil)
}

// DeleteProduct 删除产品
func (s *ProductService) DeleteProduct(id int) *dto.Result {
	// 检查产品是否存在
	existingProduct, err := s.ProductRepo.GetByID(id)
	if err != nil {
		log.Println("查询产品失败:", err)
		return errorResult(500, "系统错误")
	}

	if existingProduct == nil {
		return errorResult(404, "产品不存在")
	}

	// 删除产品
	err = s.ProductRepo.Delete(id)
	if err != nil {
		log.Println("删除产品失败:", err)
		return errorResult(500, "删除失败")
	}

	return successResult("删除成功", nil)
}

// GetProductByID 根据ID获取产品
func (s *ProductService) GetProductByID(id int) *dto.Result {
	product, err := s.ProductRepo.GetByID(id)
	if err != nil {
		log.Println("获取产品失败:", err)
		return errorResult(500, "系统错误")
	}

	if product == nil {
		return errorResult(404, "产品不存在")
	}

	// 直接返回数据，不显示弹窗
	return &dto.Result{
		Code: 200,
		Data: product,
	}
}

// GetAllProducts 获取所有产品
func (s *ProductService) GetAllProducts() *dto.Result {
	products, err := s.ProductRepo.FindAll()
	if err != nil {
		log.Println("获取所有产品失败:", err)
		return errorResult(500, "系统错误")
	}

	// 直接返回数据，不显示弹窗
	return &dto.Result{
		Code: 200,
		Data: products,
	}
}

// SearchProducts 搜索产品
func (s *ProductService) SearchProducts(queryDTO *dto.ProductQueryDTO) *dto.Result {
	products, err := s.ProductRepo.FindByCondition(queryDTO.Name, queryDTO.Type)
	if err != nil {
		log.Println("搜索产品失败:", err)
		return errorResult(500, "系统错误")
	}

	// 直接返回数据，不显示弹窗
	return &dto.Result{
		Code: 200,
		Data: products,
	}
}

// PageQueryProducts 分页查询产品
func (s *ProductService) PageQueryProducts(queryDTO *dto.ProductPageQueryDTO) *dto.Result {
	// 参数校验
	if queryDTO.Page <= 0 {
		queryDTO.Page = 1
	}
	if queryDTO.PageSize <= 0 {
		queryDTO.PageSize = 10
	}

	// 分页查询
	products, total, err := s.ProductRepo.PageQuery(queryDTO.Page, queryDTO.PageSize, queryDTO.Name, queryDTO.Type)
	if err != nil {
		log.Println("分页查询产品失败:", err)
		return errorResult(500, "系统错误")
	}

	// 封装分页结果
	pageResult := dto.NewPageResult(total, products, queryDTO.Page, queryDTO.PageSize)

	// 直接返回数据，不显示弹窗
	return &dto.Result{
		Code: 200,
		Data: pageResult,
	}
}

// GetProductTypes 获取所有产品类型
func (s *ProductService) GetProductTypes() *dto.Result {
	types, err := s.ProductRepo.GetProductTypes()
	if err != nil {
		log.Println("获取产品类型失败:", err)
		return errorResult(500, "系统错误")
	}

	// 直接返回数据，不显示弹窗
	return &dto.Result{
		Code: 200,
		Data: types,
	}
}
