package service

import (
	"agricultural_product_gin/model"
	"agricultural_product_gin/repository"
	"log"
)

// LogisticsService 物流服务
type LogisticsService struct {
	repo *repository.LogisticsRepository
}

// NewLogisticsService 创建物流服务
func NewLogisticsService(repo *repository.LogisticsRepository) *LogisticsService {
	return &LogisticsService{repo: repo}
}

// Save 保存物流信息
func (s *LogisticsService) Save(logistics *model.Logistics) (int, error) {
	return s.repo.Save(logistics)
}

// Update 更新物流信息
func (s *LogisticsService) Update(logistics *model.Logistics) error {
	return s.repo.Update(logistics)
}

// Delete 删除物流信息
func (s *LogisticsService) Delete(id int) error {
	return s.repo.Delete(id)
}

// GetByID 根据ID获取物流信息
func (s *LogisticsService) GetByID(id int) (*model.Logistics, error) {
	return s.repo.GetByID(id)
}

// FindAll 查找所有物流信息
func (s *LogisticsService) FindAll() ([]*model.Logistics, error) {
	return s.repo.FindAll()
}

// PageQuery 分页查询物流信息
func (s *LogisticsService) PageQuery(dto *model.LogisticsPageQueryDTO) (*model.LogisticsPageResult, error) {
	// 验证分页参数
	if dto.Page <= 0 {
		dto.Page = 1
	}

	if dto.Size <= 0 {
		dto.Size = 10
	}

	records, total, err := s.repo.PageQuery(dto)
	if err != nil {
		log.Println("分页查询物流信息失败:", err)
		return nil, err
	}

	return &model.LogisticsPageResult{
		Total:   total,
		Records: records,
	}, nil
}
