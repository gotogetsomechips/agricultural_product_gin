package dto

import "database/sql"

// ProductDTO 产品信息DTO
type ProductDTO struct {
	ID          int             `json:"pdId"`
	Name        string          `json:"pdName" binding:"required"`
	Type        string          `json:"type" binding:"required"`
	Image       string          `json:"image"`
	Description string          `json:"pdDescription"`
	UnitPrice   sql.NullFloat64 `json:"unitPrice"`
}

// ProductQueryDTO 产品查询DTO
type ProductQueryDTO struct {
	Name string `form:"name"`
	Type string `form:"type"`
}
