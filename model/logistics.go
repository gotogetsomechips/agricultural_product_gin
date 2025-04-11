package model

import "time"

// Logistics 物流信息模型
type Logistics struct {
	ID            int        `json:"logId"`
	ProductInfoID int        `json:"productInfoId"`
	CompanyID     int        `json:"companyId"`
	StartLocation string     `json:"startLocation"`
	Destination   string     `json:"destination"`
	StartTime     time.Time  `json:"startTime"`
	EndTime       *time.Time `json:"endTime"`

	// 关联信息 (用于展示)
	ProductName   string `json:"pdName,omitempty"`
	CompanyName   string `json:"comName,omitempty"`
	Administrator string `json:"comAdministrator,omitempty"`
	Phone         string `json:"comPhone,omitempty"`
}

// LogisticsPageQueryDTO 物流分页查询DTO
type LogisticsPageQueryDTO struct {
	Page          int    `json:"page"`
	Size          int    `json:"size"`
	LogisticsId   int    `json:"logId"`
	ProductName   string `json:"pdName"`
	CompanyName   string `json:"comName"`
	StartLocation string `json:"startLocation"`
	Destination   string `json:"destination"`
	Administrator string `json:"comAdministrator"`
	StartTime     string `json:"startTime"`
}

// LogisticsPageResult 物流分页查询结果
type LogisticsPageResult struct {
	Total   int64        `json:"total"`
	Records []*Logistics `json:"records"`
}
