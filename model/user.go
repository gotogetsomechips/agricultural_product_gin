package model

import "database/sql"

// 用户实体
type User struct {
	ID       int            `json:"id"`
	Username string         `json:"username"`
	Password string         `json:"password,omitempty"`
	Sex      sql.NullString `json:"sex"`   // 使用NullString处理NULL值
	Name     sql.NullString `json:"name"`  // 如果可能为NULL，也使用NullString
	Phone    sql.NullString `json:"phone"` // 如果可能为NULL，也使用NullString
}
