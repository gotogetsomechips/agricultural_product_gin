package model

import "time"

// ProductionInfo 生产信息实体
type ProductionInfo struct {
	ID             int       `json:"piId"`            // 生产信息ID
	ProductID      int       `json:"productId"`       // 产品ID
	ProductPlaceID int       `json:"productPlaceId"`  // 生产地ID
	SeedSource     string    `json:"seed"`            // 种子来源
	Description    string    `json:"piDescription"`   // 生产描述
	PlantingDate   time.Time `json:"plantingDate"`    // 播种时间
	HarvestDate    time.Time `json:"harvestDate"`     // 收获时间
	Administrator  string    `json:"ppAdministrator"` // 负责人
	Phone          string    `json:"ppPhone"`         // 联系方式
}

// ProductionInfoWithDetails 包含详细信息的扩展生产信息
type ProductionInfoWithDetails struct {
	ProductionInfo
	ProductName     string `json:"pdName"`    // 产品名称
	ProductionPlace string `json:"ppAddress"` // 生产地地址
}
