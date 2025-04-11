package service

import (
	"log"

	"agricultural_product_gin/dto"
	"agricultural_product_gin/model"
	"agricultural_product_gin/repository"
)

// SalePlaceService 销售地服务
type SalePlaceService struct {
	SalePlaceRepo *repository.SalePlaceRepository
}

// NewSalePlaceService 创建销售地服务
func NewSalePlaceService(salePlaceRepo *repository.SalePlaceRepository) *SalePlaceService {
	return &SalePlaceService{SalePlaceRepo: salePlaceRepo}
}

// CreateSalePlace 创建销售地
func (s *SalePlaceService) CreateSalePlace(salePlaceDTO *dto.SalePlaceDTO) *dto.Result {
	// 转换DTO为模型
	salePlace := &model.SalePlace{
		Address:       salePlaceDTO.Address,
		Administrator: salePlaceDTO.Administrator,
		Phone:         salePlaceDTO.Phone,
	}

	// 保存销售地
	id, err := s.SalePlaceRepo.Save(salePlace)
	if err != nil {
		log.Println("创建销售地失败:", err)
		return errorResult(500, "创建销售地失败")
	}

	// 返回创建成功的销售地ID
	return successResult("添加成功", id)
}

// UpdateSalePlace 更新销售地
func (s *SalePlaceService) UpdateSalePlace(salePlaceDTO *dto.SalePlaceDTO) *dto.Result {
	// 检查销售地是否存在
	existingSalePlace, err := s.SalePlaceRepo.GetByID(salePlaceDTO.ID)
	if err != nil {
		log.Println("查询销售地失败:", err)
		return errorResult(500, "系统错误")
	}

	if existingSalePlace == nil {
		return errorResult(404, "销售地不存在")
	}

	// 转换DTO为模型
	salePlace := &model.SalePlace{
		ID:            salePlaceDTO.ID,
		Address:       salePlaceDTO.Address,
		Administrator: salePlaceDTO.Administrator,
		Phone:         salePlaceDTO.Phone,
	}

	// 更新销售地
	err = s.SalePlaceRepo.Update(salePlace)
	if err != nil {
		log.Println("更新销售地失败:", err)
		return errorResult(500, "更新失败")
	}

	return successResult("更新成功", nil)
}

// DeleteSalePlace 删除销售地
func (s *SalePlaceService) DeleteSalePlace(id int) *dto.Result {
	// 检查销售地是否存在
	existingSalePlace, err := s.SalePlaceRepo.GetByID(id)
	if err != nil {
		log.Println("查询销售地失败:", err)
		return errorResult(500, "系统错误")
	}

	if existingSalePlace == nil {
		return errorResult(404, "销售地不存在")
	}

	// 删除销售地
	err = s.SalePlaceRepo.Delete(id)
	if err != nil {
		log.Println("删除销售地失败:", err)
		return errorResult(500, "删除失败")
	}

	return successResult("删除成功", nil)
}

// GetSalePlaceByID 根据ID获取销售地
func (s *SalePlaceService) GetSalePlaceByID(id int) *dto.Result {
	salePlace, err := s.SalePlaceRepo.GetByID(id)
	if err != nil {
		log.Println("获取销售地失败:", err)
		return errorResult(500, "系统错误")
	}

	if salePlace == nil {
		return errorResult(404, "销售地不存在")
	}

	// 直接返回数据，不显示弹窗
	return &dto.Result{
		Code: 200,
		Data: salePlace,
	}
}

// GetAllSalePlaces 获取所有销售地
func (s *SalePlaceService) GetAllSalePlaces() *dto.Result {
	salePlaces, err := s.SalePlaceRepo.FindAll()
	if err != nil {
		log.Println("获取所有销售地失败:", err)
		return errorResult(500, "系统错误")
	}

	// 直接返回数据，不显示弹窗
	return &dto.Result{
		Code: 200,
		Data: salePlaces,
	}
}

// PageQuerySalePlaces 分页查询销售地
func (s *SalePlaceService) PageQuerySalePlaces(queryDTO *dto.SalePlacePageQueryDTO) *dto.Result {
	// 参数校验
	if queryDTO.Page <= 0 {
		queryDTO.Page = 1
	}
	if queryDTO.PageSize <= 0 {
		queryDTO.PageSize = 10
	}

	// 分页查询
	salePlaces, total, err := s.SalePlaceRepo.PageQuery(
		queryDTO.Page,
		queryDTO.PageSize,
		queryDTO.ID,
		queryDTO.Address,
		queryDTO.Administrator,
		queryDTO.Phone,
	)
	if err != nil {
		log.Println("分页查询销售地失败:", err)
		return errorResult(500, "系统错误")
	}

	// 封装分页结果
	pageResult := dto.NewPageResult(total, salePlaces, queryDTO.Page, queryDTO.PageSize)

	// 直接返回数据，不显示弹窗
	return &dto.Result{
		Code: 200,
		Data: pageResult,
	}
}
