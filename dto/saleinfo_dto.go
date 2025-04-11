package dto

import "time"

// SaleInfoDTO 销售信息DTO
type SaleInfoDTO struct {
	ID          int       `json:"siId"`
	LogisticsID int       `json:"logisticsId" binding:"required"`
	SalePlaceID int       `json:"salePlaceId" binding:"required"`
	Description string    `json:"siDescription"`
	SaleTime    time.Time `json:"saleTime" binding:"required"`
}

// SaleInfoPageQueryDTO 销售信息分页查询DTO
type SaleInfoPageQueryDTO struct {
	Page        int       `json:"page" binding:"required"`
	Size        int       `json:"size" binding:"required"`
	SaleInfoID  int       `json:"saleInfoId"`
	ProductName string    `json:"productName"`
	SalePlace   string    `json:"salePlace"`
	SaleTime    time.Time `json:"saleTime"`
}
