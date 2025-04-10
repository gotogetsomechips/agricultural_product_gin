package model

// ProductionPlace 生产地实体
type ProductionPlace struct {
	ID            int    `json:"ppId"`            // 生产地ID
	Address       string `json:"ppAddress"`       // 生产地地址
	Administrator string `json:"ppAdministrator"` // 负责人
	Phone         string `json:"ppPhone"`         // 联系电话
}
