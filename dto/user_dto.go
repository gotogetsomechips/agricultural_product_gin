package dto

// UserRegAndLoginDTO 用户注册和登录DTO
type UserRegAndLoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserDTO 用户信息编辑DTO
type UserDTO struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Sex      string `json:"sex"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
}

// UserEditPasswordDTO 密码修改DTO
type UserEditPasswordDTO struct {
	OldPassword     string `json:"oldPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

// 返回结果结构
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResult(code int, msg string, data interface{}) *Result {
	return &Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
