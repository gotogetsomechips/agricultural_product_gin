package dto

// ProductionPlaceDTO 生产地信息DTO
type ProductionPlaceDTO struct {
	ID            int    `json:"ppId"`
	Address       string `json:"ppAddress" binding:"required"`
	Administrator string `json:"ppAdministrator" binding:"required"`
	Phone         string `json:"ppPhone" binding:"required"`
}

// ProductionPlacePageQueryDTO 生产地分页查询DTO
type ProductionPlacePageQueryDTO struct {
	Page          int    `json:"page"`
	PageSize      int    `json:"size"`
	ID            string `json:"ppId"`            // 生产地编号
	Address       string `json:"ppAddress"`       // 生产地地址
	Administrator string `json:"ppAdministrator"` // 负责人
}
