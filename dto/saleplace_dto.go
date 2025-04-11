package dto

// SalePlaceDTO 销售地信息DTO
type SalePlaceDTO struct {
	ID            int    `json:"spId"`
	Address       string `json:"spAddress" binding:"required"`
	Administrator string `json:"spAdministrator"`
	Phone         string `json:"spPhone"`
}

// SalePlacePageQueryDTO 销售地分页查询DTO
type SalePlacePageQueryDTO struct {
	Page          int    `json:"page" binding:"required"`
	PageSize      int    `json:"size" binding:"required"`
	ID            string `json:"spId"`
	Address       string `json:"spAddress"`
	Administrator string `json:"spAdministrator"`
	Phone         string `json:"spPhone"`
}
