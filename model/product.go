package model

// Product 产品实体
type Product struct {
	ID          int    `json:"pdId"`
	Name        string `json:"pdName"`
	Type        string `json:"type"`
	Image       string `json:"image"`
	Description string `json:"pdDescription"`
}
