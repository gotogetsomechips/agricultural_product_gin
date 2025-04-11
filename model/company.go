package model

// Company 公司实体
type Company struct {
	ID            int    `json:"comId"`
	Name          string `json:"comName"`
	Address       string `json:"comAddress"`
	Administrator string `json:"comAdministrator"`
	Phone         string `json:"comPhone"`
}
