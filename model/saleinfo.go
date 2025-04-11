package model

import "time"

// SaleInfo 销售信息模型
type SaleInfo struct {
	ID          int       `json:"siId"`          // 销售ID
	LogisticsID int       `json:"logisticsId"`   // 物流ID
	SalePlaceID int       `json:"salePlaceId"`   // 销售地ID
	Description string    `json:"siDescription"` // 说明
	SaleTime    time.Time `json:"saleTime"`      // 销售时间
}

// SaleInfoVO 销售信息视图对象(包含关联信息)
type SaleInfoVO struct {
	ID            int       `json:"siId"`
	LogisticsID   int       `json:"logisticsId"`
	SalePlaceID   int       `json:"salePlaceId"`
	Description   string    `json:"siDescription"`
	SaleTime      time.Time `json:"saleTime"`
	ProductName   string    `json:"pdName"`          // 产品名称
	SalePlace     string    `json:"spAddress"`       // 销售地地址
	Administrator string    `json:"spAdministrator"` // 销售地负责人
	StartLocation string    `json:"startLocation"`   // 物流起始地
	Destination   string    `json:"destination"`     // 物流目的地
}

// SaleInfoPageQuery 分页查询参数
type SaleInfoPageQuery struct {
	Page        int       `json:"page"`        // 页码
	Size        int       `json:"size"`        // 每页大小
	SaleInfoID  int       `json:"saleInfoId"`  // 销售ID
	ProductName string    `json:"productName"` // 产品名称
	SalePlace   string    `json:"salePlace"`   // 销售地
	SaleTime    time.Time `json:"saleTime"`    // 销售时间
}
