package model

// PageResult 分页结果
type PageResult struct {
	Total   int64       `json:"total"`   // 总记录数
	Records interface{} `json:"records"` // 当前页数据
}
