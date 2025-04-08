package service

import (
	"database/sql"
	"errors"
	"log"

	"agricultural_product_gin/dto"
	"agricultural_product_gin/model"
	"agricultural_product_gin/repository"
	"agricultural_product_gin/utils"
)

// 错误信息常量
const (
	UsernameError       = "用户名已被占用"
	UsernameInvalid     = "用户名不存在"
	PasswordInvalid     = "密码错误"
	PasswordEditInvalid = "密码参数不完整"
	PasswordError       = "两次输入的密码不一致"
)

// UserService 用户服务
type UserService struct {
	UserRepo *repository.UserRepository
}

// 辅助函数
func successResult(msg string, data interface{}) *dto.Result {
	return &dto.Result{Code: 200, Msg: msg, Data: data}
}

func errorResult(code int, msg string) *dto.Result {
	return &dto.Result{Code: code, Msg: msg, Data: nil}
}

// NewUserService 创建用户服务
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

// Register 用户注册
func (s *UserService) Register(dto *dto.UserRegAndLoginDTO) *dto.Result {
	username := dto.Username
	password := dto.Password

	// 检查用户名是否已存在
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		log.Println("查询用户失败:", err)
		return errorResult(500, "系统错误")
	}

	if user != nil {
		// 用户名已被占用
		return errorResult(400, UsernameError)
	}

	// 加密密码
	encryptedPassword := utils.EncryptPassword(password)

	// 保存用户
	err = s.UserRepo.Save(username, encryptedPassword)
	if err != nil {
		log.Println("保存用户失败:", err)
		return errorResult(500, "注册失败")
	}

	return successResult("注册成功", nil)
}

// Login 用户登录
func (s *UserService) Login(dto *dto.UserRegAndLoginDTO) *dto.Result {
	username := dto.Username
	password := dto.Password

	// 查询用户
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		log.Println("查询用户失败:", err)
		return errorResult(500, "系统错误")
	}

	if user == nil {
		return errorResult(400, UsernameInvalid)
	}

	// 验证密码
	encryptedPassword := utils.EncryptPassword(password)
	if encryptedPassword != user.Password {
		return errorResult(400, PasswordInvalid)
	}

	// 生成Token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		log.Println("生成Token失败:", err)
		return errorResult(500, "登录失败")
	}

	return successResult("登录成功", token)
}

// Logout 退出登录
func (s *UserService) Logout() *dto.Result {
	// 清除ThreadLocal数据
	threadLocal := utils.GetUserLocal()
	threadLocal.Remove("userID")
	threadLocal.Remove("username")

	return successResult("退出成功", nil)
}

// GetUserInfo 获取用户信息
func (s *UserService) GetUserInfo() (*model.User, error) {
	threadLocal := utils.GetUserLocal()
	userID, ok := threadLocal.Get("userID").(int)
	if !ok {
		return nil, errors.New("未登录")
	}

	user, err := s.UserRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	// 隐藏密码
	user.Password = "******"
	return user, nil
}

// Update 更新用户信息
func (s *UserService) Update(userDTO *dto.UserDTO) *dto.Result {
	// 检查用户名是否已存在（但排除当前用户）
	existingUser, err := s.UserRepo.FindByUsername(userDTO.Username)
	if err != nil {
		log.Println("查询用户失败:", err)
		return errorResult(500, "系统错误")
	}

	if existingUser != nil && existingUser.ID != userDTO.ID {
		return errorResult(400, UsernameError)
	}

	// 更新用户信息
	user := &model.User{
		ID:       userDTO.ID,
		Username: userDTO.Username,
		Sex:      sql.NullString{String: userDTO.Sex, Valid: userDTO.Sex != ""},
		Name:     sql.NullString{String: userDTO.Name, Valid: userDTO.Name != ""},
		Phone:    sql.NullString{String: userDTO.Phone, Valid: userDTO.Phone != ""},
	}
	
	err = s.UserRepo.Update(user)
	if err != nil {
		log.Println("更新用户失败:", err)
		return errorResult(500, "更新失败")
	}

	return successResult("更新成功", nil)
}

// EditPassword 修改密码
func (s *UserService) EditPassword(dto *dto.UserEditPasswordDTO) *dto.Result {
	oldPassword := dto.OldPassword
	newPassword := dto.NewPassword
	confirmPassword := dto.ConfirmPassword

	// 参数校验
	if oldPassword == "" || newPassword == "" || confirmPassword == "" {
		return errorResult(400, PasswordEditInvalid)
	}

	// 获取当前用户
	threadLocal := utils.GetUserLocal()
	username, ok := threadLocal.Get("username").(string)
	if !ok {
		return errorResult(401, "未登录")
	}

	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		log.Println("获取用户失败:", err)
		return errorResult(500, "系统错误")
	}

	// 验证旧密码
	encryptedOldPassword := utils.EncryptPassword(oldPassword)
	if encryptedOldPassword != user.Password {
		return errorResult(400, PasswordInvalid)
	}

	// 检查两次密码是否一致
	if newPassword != confirmPassword {
		return errorResult(400, PasswordError)
	}

	// 更新密码
	encryptedNewPassword := utils.EncryptPassword(newPassword)
	err = s.UserRepo.UpdatePassword(user.ID, encryptedNewPassword)
	if err != nil {
		log.Println("更新密码失败:", err)
		return errorResult(500, "修改密码失败")
	}

	return successResult("修改密码成功", newPassword)
}
