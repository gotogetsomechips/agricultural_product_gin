package controller

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UploadController 处理文件上传相关逻辑
type UploadController struct {
	uploadPath string
}

// NewUploadController 创建一个新的上传控制器
func NewUploadController() *UploadController {
	// 获取项目根目录
	projectRoot, _ := os.Getwd()

	// 设置上传路径 - 与Java版本保持一致
	uploadPath := filepath.Join(projectRoot, "resources", "images")

	// 确保目录存在
	if err := os.MkdirAll(uploadPath, 0755); err != nil {
		log.Printf("创建上传目录失败: %v", err)
	}

	log.Printf("文件存储路径：%s", uploadPath)

	return &UploadController{
		uploadPath: uploadPath,
	}
}

// Result 定义API返回结构，与你项目中已有的结构保持一致
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Upload 处理文件上传请求
func (uc *UploadController) Upload(c *gin.Context) {
	// 检查认证头 - 如果需要认证的话
	// 假设你的middleware.JWTMiddleware已经验证了token

	// 从请求中获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, Result{
			Code: 400,
			Msg:  "获取文件失败: " + err.Error(),
		})
		return
	}

	// 生成安全的文件名（带UUID）
	fileExt := filepath.Ext(file.Filename)
	fileName := uuid.New().String() + fileExt

	// 构建完整的文件路径
	filePath := filepath.Join(uc.uploadPath, fileName)

	// 保存文件
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		log.Printf("文件操作失败 | 路径：%s | 错误：%v", filePath, err)
		c.JSON(http.StatusOK, Result{
			Code: 500,
			Msg:  "文件操作失败: " + err.Error(),
		})
		return
	}

	// 返回相对访问路径（与Java版本对应）
	fileURL := "http://localhost:8080/images/" + fileName

	c.JSON(http.StatusOK, Result{
		Code: 200,
		Msg:  "上传成功",
		Data: fileURL,
	})
}

func (uc *UploadController) RegisterStaticRoutes(router *gin.Engine) {
	// 设置静态文件服务，用于访问上传的图片
	router.StaticFS("/images", http.Dir(uc.uploadPath))
}
