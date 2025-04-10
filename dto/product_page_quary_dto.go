package dto

// ProductPageQueryDTO 产品分页查询DTO
type ProductPageQueryDTO struct {
	Page     int    `json:"page"`
	PageSize int    `json:"size"`
	Name     string `json:"productName"`
	Type     string `json:"type"`
}

// PageResult 分页结果
type PageResult struct {
	Total    int64       `json:"total"`    // 总记录数
	Records  interface{} `json:"records"`  // 当前页数据
	Page     int         `json:"page"`     // 当前页码
	PageSize int         `json:"pageSize"` // 每页记录数
}

// NewPageResult 创建分页结果
func NewPageResult(total int64, records interface{}, page, pageSize int) *PageResult {
	return &PageResult{
		Total:    total,
		Records:  records,
		Page:     page,
		PageSize: pageSize,
	}
}
