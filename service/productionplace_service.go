package service

import (
	"log"

	"agricultural_product_gin/dto"
	"agricultural_product_gin/model"
	"agricultural_product_gin/repository"
)

// ProductionPlaceService 生产地信息服务
type ProductionPlaceService struct {
	ProductionPlaceRepo *repository.ProductionPlaceRepository
}

// NewProductionPlaceService 创建生产地信息服务
func NewProductionPlaceService(repo *repository.ProductionPlaceRepository) *ProductionPlaceService {
	return &ProductionPlaceService{ProductionPlaceRepo: repo}
}

// CreateProductionPlace 创建生产地信息
func (s *ProductionPlaceService) CreateProductionPlace(dto *dto.ProductionPlaceDTO) *dto.Result {
	// 转换DTO为模型
	place := &model.ProductionPlace{
		Address:       dto.Address,
		Administrator: dto.Administrator,
		Phone:         dto.Phone,
	}

	// 保存生产地信息
	id, err := s.ProductionPlaceRepo.Save(place)
	if err != nil {
		log.Println("创建生产地信息失败:", err)
		return errorResult(500, "创建生产地信息失败")
	}

	return successResult("添加成功", id)
}

// UpdateProductionPlace 更新生产地信息
func (s *ProductionPlaceService) UpdateProductionPlace(dto *dto.ProductionPlaceDTO) *dto.Result {
	// 检查生产地信息是否存在
	existing, err := s.ProductionPlaceRepo.GetByID(dto.ID)
	if err != nil {
		log.Println("查询生产地信息失败:", err)
		return errorResult(500, "系统错误")
	}
	if existing == nil {
		return errorResult(404, "生产地信息不存在")
	}

	// 转换DTO为模型
	place := &model.ProductionPlace{
		ID:            dto.ID,
		Address:       dto.Address,
		Administrator: dto.Administrator,
		Phone:         dto.Phone,
	}

	// 更新生产地信息
	err = s.ProductionPlaceRepo.Update(place)
	if err != nil {
		log.Println("更新生产地信息失败:", err)
		return errorResult(500, "更新失败")
	}

	return successResult("更新成功", nil)
}

// DeleteProductionPlace 删除生产地信息
func (s *ProductionPlaceService) DeleteProductionPlace(id int) *dto.Result {
	// 检查生产地信息是否存在
	existing, err := s.ProductionPlaceRepo.GetByID(id)
	if err != nil {
		log.Println("查询生产地信息失败:", err)
		return errorResult(500, "系统错误")
	}
	if existing == nil {
		return errorResult(404, "生产地信息不存在")
	}

	// 删除生产地信息
	err = s.ProductionPlaceRepo.Delete(id)
	if err != nil {
		log.Println("删除生产地信息失败:", err)
		return errorResult(500, "删除失败")
	}

	return successResult("删除成功", nil)
}

// GetProductionPlaceByID 根据ID获取生产地信息
func (s *ProductionPlaceService) GetProductionPlaceByID(id int) *dto.Result {
	place, err := s.ProductionPlaceRepo.GetByID(id)
	if err != nil {
		log.Println("获取生产地信息失败:", err)
		return errorResult(500, "系统错误")
	}

	if place == nil {
		return errorResult(404, "生产地信息不存在")
	}

	return &dto.Result{
		Code: 200,
		Data: place,
	}
}

// PageQueryProductionPlaces 分页查询生产地信息
func (s *ProductionPlaceService) PageQueryProductionPlaces(queryDTO *dto.ProductionPlacePageQueryDTO) *dto.Result {
	// 参数校验
	if queryDTO.Page <= 0 {
		queryDTO.Page = 1
	}
	if queryDTO.PageSize <= 0 {
		queryDTO.PageSize = 10
	}

	// 分页查询
	places, total, err := s.ProductionPlaceRepo.PageQuery(
		queryDTO.Page, queryDTO.PageSize,
		queryDTO.ID, queryDTO.Address, queryDTO.Administrator)
	if err != nil {
		log.Println("分页查询生产地信息失败:", err)
		return errorResult(500, "系统错误")
	}

	// 封装分页结果
	pageResult := dto.NewPageResult(total, places, queryDTO.Page, queryDTO.PageSize)

	return &dto.Result{
		Code: 200,
		Data: pageResult,
	}
}

// GetAllProductionPlaces 获取所有生产地信息
func (s *ProductionPlaceService) GetAllProductionPlaces() *dto.Result {
	places, err := s.ProductionPlaceRepo.GetAll()
	if err != nil {
		log.Println("获取所有生产地信息失败:", err)
		return errorResult(500, "系统错误")
	}

	return &dto.Result{
		Code: 200,
		Data: places,
	}
}
