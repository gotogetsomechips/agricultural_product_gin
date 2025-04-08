package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"agricultural_product_gin/dto"
	"agricultural_product_gin/service"
)

// UserController 用户控制器
type UserController struct {
	UserService *service.UserService
}

// NewUserController 创建用户控制器
func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: userService}
}

// Register 注册
func (c *UserController) Register(ctx *gin.Context) {
	var userDTO dto.UserRegAndLoginDTO
	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	log.Printf("用户名：%s，密码：%s", userDTO.Username, userDTO.Password)
	result := c.UserService.Register(&userDTO)
	ctx.JSON(http.StatusOK, result)
}

// Login 登录
func (c *UserController) Login(ctx *gin.Context) {
	var userDTO dto.UserRegAndLoginDTO
	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	log.Printf("用户名：%s，密码：%s", userDTO.Username, userDTO.Password)
	result := c.UserService.Login(&userDTO)
	ctx.JSON(http.StatusOK, result)
}

// Logout 退出登录
func (c *UserController) Logout(ctx *gin.Context) {
	result := c.UserService.Logout()
	ctx.JSON(http.StatusOK, result)
}

// GetUserInfo 获取用户信息
func (c *UserController) GetUserInfo(ctx *gin.Context) {
	user, err := c.UserService.GetUserInfo()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": user,
	})
}

// Update 更新用户信息
func (c *UserController) Update(ctx *gin.Context) {
	var userDTO dto.UserDTO
	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	log.Printf("编辑用户信息：%+v", userDTO)
	result := c.UserService.Update(&userDTO)
	ctx.JSON(http.StatusOK, result)
}

// EditPassword 修改密码
func (c *UserController) EditPassword(ctx *gin.Context) {
	var passwordDTO dto.UserEditPasswordDTO
	if err := ctx.ShouldBindJSON(&passwordDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	log.Printf("修改密码：%+v", passwordDTO)
	result := c.UserService.EditPassword(&passwordDTO)
	ctx.JSON(http.StatusOK, result)
}
