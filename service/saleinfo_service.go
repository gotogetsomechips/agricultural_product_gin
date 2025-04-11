package service

import (
	"log"

	"agricultural_product_gin/dto"
	"agricultural_product_gin/model"
	"agricultural_product_gin/repository"
)

// SaleInfoService 销售信息服务接口
type SaleInfoService interface {
	Save(saleInfoDTO *dto.SaleInfoDTO) *dto.Result
	Update(saleInfoDTO *dto.SaleInfoDTO) *dto.Result
	Delete(id int) *dto.Result
	GetByID(id int) *dto.Result
	GetAll() *dto.Result
	PageQuery(queryDTO *dto.SaleInfoPageQueryDTO) *dto.Result
}

// SaleInfoServiceImpl 销售信息服务实现
type SaleInfoServiceImpl struct {
	repo *repository.SaleInfoRepository // 改为指针类型
}

func NewSaleInfoService(repo *repository.SaleInfoRepository) SaleInfoService {
	return &SaleInfoServiceImpl{repo: repo}
}

// Save 保存销售信息
func (s *SaleInfoServiceImpl) Save(saleInfoDTO *dto.SaleInfoDTO) *dto.Result {
	// 转换DTO为模型
	saleInfo := &model.SaleInfo{
		LogisticsID: saleInfoDTO.LogisticsID,
		SalePlaceID: saleInfoDTO.SalePlaceID,
		Description: saleInfoDTO.Description,
		SaleTime:    saleInfoDTO.SaleTime,
	}

	id, err := s.repo.Save(saleInfo)
	if err != nil {
		log.Println("保存销售信息失败:", err)
		return errorResult(500, "保存失败")
	}

	return successResult("保存成功", id)
}

// Update 更新销售信息
func (s *SaleInfoServiceImpl) Update(saleInfoDTO *dto.SaleInfoDTO) *dto.Result {
	// 检查销售信息是否存在
	existingSaleInfo, err := s.repo.GetByID(saleInfoDTO.ID)
	if err != nil {
		log.Println("查询销售信息失败:", err)
		return errorResult(500, "系统错误")
	}

	if existingSaleInfo == nil {
		return errorResult(404, "销售信息不存在")
	}

	// 转换DTO为模型
	saleInfo := &model.SaleInfo{
		ID:          saleInfoDTO.ID,
		LogisticsID: saleInfoDTO.LogisticsID,
		SalePlaceID: saleInfoDTO.SalePlaceID,
		Description: saleInfoDTO.Description,
		SaleTime:    saleInfoDTO.SaleTime,
	}

	// 更新销售信息
	err = s.repo.Update(saleInfo)
	if err != nil {
		log.Println("更新销售信息失败:", err)
		return errorResult(500, "更新失败")
	}

	return successResult("更新成功", nil)
}

// Delete 删除销售信息
func (s *SaleInfoServiceImpl) Delete(id int) *dto.Result {
	// 检查销售信息是否存在
	existingSaleInfo, err := s.repo.GetByID(id)
	if err != nil {
		log.Println("查询销售信息失败:", err)
		return errorResult(500, "系统错误")
	}

	if existingSaleInfo == nil {
		return errorResult(404, "销售信息不存在")
	}

	// 删除销售信息
	err = s.repo.Delete(id)
	if err != nil {
		log.Println("删除销售信息失败:", err)
		return errorResult(500, "删除失败")
	}

	return successResult("删除成功", nil)
}

// GetByID 根据ID获取销售信息
func (s *SaleInfoServiceImpl) GetByID(id int) *dto.Result {
	saleInfo, err := s.repo.GetByID(id)
	if err != nil {
		log.Println("获取销售信息失败:", err)
		return errorResult(500, "系统错误")
	}

	if saleInfo == nil {
		return errorResult(404, "销售信息不存在")
	}

	return successResult("查询成功", saleInfo)
}

// GetAll 获取所有销售信息
func (s *SaleInfoServiceImpl) GetAll() *dto.Result {
	saleInfos, err := s.repo.FindAll()
	if err != nil {
		log.Println("获取所有销售信息失败:", err)
		return errorResult(500, "系统错误")
	}

	return successResult("查询成功", saleInfos)
}

// PageQuery 分页查询销售信息
func (s *SaleInfoServiceImpl) PageQuery(queryDTO *dto.SaleInfoPageQueryDTO) *dto.Result {
	// 参数校验
	if queryDTO.Page <= 0 {
		queryDTO.Page = 1
	}
	if queryDTO.Size <= 0 {
		queryDTO.Size = 10
	}

	// 转换DTO为模型
	query := &model.SaleInfoPageQuery{
		Page:        queryDTO.Page,
		Size:        queryDTO.Size,
		SaleInfoID:  queryDTO.SaleInfoID,
		ProductName: queryDTO.ProductName,
		SalePlace:   queryDTO.SalePlace,
		SaleTime:    queryDTO.SaleTime,
	}

	// 分页查询
	saleInfos, total, err := s.repo.PageQuery(query)
	if err != nil {
		log.Println("分页查询销售信息失败:", err)
		return errorResult(500, "系统错误")
	}

	// 封装分页结果
	pageResult := dto.NewPageResult(total, saleInfos, queryDTO.Page, queryDTO.Size)

	return successResult("查询成功", pageResult)
}
