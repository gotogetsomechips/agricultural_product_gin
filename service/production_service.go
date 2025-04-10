package service

import (
	"log"

	"agricultural_product_gin/dto"
	"agricultural_product_gin/model"
	"agricultural_product_gin/repository"
)

// ProductionService 生产信息服务
type ProductionService struct {
	ProductionRepo *repository.ProductionRepository
}

// NewProductionService 创建生产信息服务
func NewProductionService(repo *repository.ProductionRepository) *ProductionService {
	return &ProductionService{ProductionRepo: repo}
}

// CreateProduction 创建生产信息
func (s *ProductionService) CreateProduction(dto *dto.ProductionDTO) *dto.Result {
	// 转换DTO为模型
	production := &model.ProductionInfo{
		ProductID:      dto.ProductID,
		ProductPlaceID: dto.ProductPlaceID,
		SeedSource:     dto.SeedSource,
		Description:    dto.Description,
		PlantingDate:   dto.PlantingDate,
		HarvestDate:    dto.HarvestDate,
	}

	// 保存生产信息
	id, err := s.ProductionRepo.Save(production)
	if err != nil {
		log.Println("创建生产信息失败:", err)
		return errorResult(500, "创建生产信息失败")
	}

	return successResult("添加成功", id)
}

// UpdateProduction 更新生产信息
func (s *ProductionService) UpdateProduction(dto *dto.ProductionDTO) *dto.Result {
	// 检查生产信息是否存在
	existing, err := s.ProductionRepo.GetByID(dto.ID)
	if err != nil {
		log.Println("查询生产信息失败:", err)
		return errorResult(500, "系统错误")
	}
	if existing == nil {
		return errorResult(404, "生产信息不存在")
	}

	// 转换DTO为模型
	production := &model.ProductionInfo{
		ID:             dto.ID,
		ProductID:      dto.ProductID,
		ProductPlaceID: dto.ProductPlaceID,
		SeedSource:     dto.SeedSource,
		Description:    dto.Description,
		PlantingDate:   dto.PlantingDate,
		HarvestDate:    dto.HarvestDate,
	}

	// 更新生产信息
	err = s.ProductionRepo.Update(production)
	if err != nil {
		log.Println("更新生产信息失败:", err)
		return errorResult(500, "更新失败")
	}

	return successResult("更新成功", nil)
}

// DeleteProduction 删除生产信息
func (s *ProductionService) DeleteProduction(id int) *dto.Result {
	// 检查生产信息是否存在
	existing, err := s.ProductionRepo.GetByID(id)
	if err != nil {
		log.Println("查询生产信息失败:", err)
		return errorResult(500, "系统错误")
	}
	if existing == nil {
		return errorResult(404, "生产信息不存在")
	}

	// 删除生产信息
	err = s.ProductionRepo.Delete(id)
	if err != nil {
		log.Println("删除生产信息失败:", err)
		return errorResult(500, "删除失败")
	}

	return successResult("删除成功", nil)
}

// GetProductionByID 根据ID获取生产信息
func (s *ProductionService) GetProductionByID(id int) *dto.Result {
	production, err := s.ProductionRepo.GetByID(id)
	if err != nil {
		log.Println("获取生产信息失败:", err)
		return errorResult(500, "系统错误")
	}

	if production == nil {
		return errorResult(404, "生产信息不存在")
	}

	return &dto.Result{
		Code: 200,
		Data: production,
	}
}

// PageQueryProductions 分页查询生产信息
func (s *ProductionService) PageQueryProductions(queryDTO *dto.ProductionPageQueryDTO) *dto.Result {
	// 参数校验
	if queryDTO.Page <= 0 {
		queryDTO.Page = 1
	}
	if queryDTO.PageSize <= 0 {
		queryDTO.PageSize = 10
	}

	// 分页查询
	productions, total, err := s.ProductionRepo.PageQuery(
		queryDTO.Page, queryDTO.PageSize,
		queryDTO.ProductInfoID, queryDTO.ProductName,
		queryDTO.ProductPlace, queryDTO.Seed, queryDTO.Administrator)
	if err != nil {
		log.Println("分页查询生产信息失败:", err)
		return errorResult(500, "系统错误")
	}

	// 封装分页结果
	pageResult := dto.NewPageResult(total, productions, queryDTO.Page, queryDTO.PageSize)

	return &dto.Result{
		Code: 200,
		Data: pageResult,
	}
}

// GetAllProductions 获取所有生产信息
func (s *ProductionService) GetAllProductions() *dto.Result {
	productions, err := s.ProductionRepo.GetAll()
	if err != nil {
		log.Println("获取所有生产信息失败:", err)
		return errorResult(500, "系统错误")
	}

	return &dto.Result{
		Code: 200,
		Data: productions,
	}
}
