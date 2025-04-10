package dto

import "time"

// ProductionDTO 生产信息DTO
type ProductionDTO struct {
	ID             int       `json:"piId"`
	ProductID      int       `json:"productId" binding:"required"`
	ProductPlaceID int       `json:"productPlaceId" binding:"required"`
	SeedSource     string    `json:"seed" binding:"required"`
	Description    string    `json:"piDescription"`
	PlantingDate   time.Time `json:"plantingDate" binding:"required"`
	HarvestDate    time.Time `json:"harvestDate" binding:"required"`
}

// ProductionPageQueryDTO 生产信息分页查询DTO
type ProductionPageQueryDTO struct {
	Page          int    `json:"page"`
	PageSize      int    `json:"size"`
	ProductInfoID string `json:"productInfoId"` // 生产编号
	ProductName   string `json:"productName"`   // 产品名称
	ProductPlace  string `json:"productPlace"`  // 生产地
	Seed          string `json:"seed"`          // 种子来源
	Administrator string `json:"administrator"` // 负责人
}
