package model

// SalePlace 销售地实体
type SalePlace struct {
	ID            int    `json:"spId"`
	Address       string `json:"spAddress"`
	Administrator string `json:"spAdministrator"`
	Phone         string `json:"spPhone"`
}
