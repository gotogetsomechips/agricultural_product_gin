package dto

// CompanyDTO 公司信息DTO
type CompanyDTO struct {
	ID            int    `json:"comId"`
	Name          string `json:"comName" binding:"required"`
	Address       string `json:"comAddress" binding:"required"`
	Administrator string `json:"comAdministrator"`
	Phone         string `json:"comPhone"`
}

// CompanyPageQueryDTO 公司分页查询DTO
type CompanyPageQueryDTO struct {
	Page          int    `json:"page" binding:"required"`
	PageSize      int    `json:"size" binding:"required"`
	Name          string `json:"comName"`
	Address       string `json:"comAddress"`
	Administrator string `json:"comAdministrator"`
	Phone         string `json:"comPhone"`
}
