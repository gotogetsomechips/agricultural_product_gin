package controller

import (
	"agricultural_product_gin/model"
	"agricultural_product_gin/service"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// LogisticsController 物流控制器
type LogisticsController struct {
	service *service.LogisticsService
}

// NewLogisticsController 创建物流控制器
func NewLogisticsController(service *service.LogisticsService) *LogisticsController {
	return &LogisticsController{service: service}
}

// Save 保存物流信息
func (c *LogisticsController) Save(ctx *gin.Context) {
	var request struct {
		ProductInfoID int    `json:"productInfoId"`
		CompanyID     int    `json:"companyId"`
		StartLocation string `json:"startLocation"`
		Destination   string `json:"destination"`
		StartTime     string `json:"startTime"`
		EndTime       string `json:"endTime"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误: " + err.Error(),
		})
		return
	}

	// 验证必要字段
	if request.ProductInfoID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "产品信息ID不能为空",
		})
		return
	}

	// 转换时间
	startTime, err := time.Parse(time.RFC3339, request.StartTime)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "开始时间格式错误",
		})
		return
	}

	var endTime *time.Time
	if request.EndTime != "" {
		et, err := time.Parse(time.RFC3339, request.EndTime)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "结束时间格式错误",
			})
			return
		}
		endTime = &et
	}

	// 构建物流对象
	logistics := &model.Logistics{
		ProductInfoID: request.ProductInfoID,
		CompanyID:     request.CompanyID,
		StartLocation: request.StartLocation,
		Destination:   request.Destination,
		StartTime:     startTime,
		EndTime:       endTime,
	}

	id, err := c.service.Save(logistics)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "保存物流信息失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "保存成功",
		"data": id,
	})
}

// Delete 删除物流信息
func (c *LogisticsController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的ID",
		})
		return
	}

	err = c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "删除物流信息失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

// GetById 根据ID获取物流信息
func (c *LogisticsController) GetById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的ID",
		})
		return
	}

	logistics, err := c.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "获取物流信息失败: " + err.Error(),
		})
		return
	}

	if logistics == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "未找到对应的物流信息",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": logistics,
	})
}

// PageQuery 分页查询物流信息
func (c *LogisticsController) PageQuery(ctx *gin.Context) {
	var dto model.LogisticsPageQueryDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		log.Println("绑定请求参数失败:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	pageResult, err := c.service.PageQuery(&dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "查询物流信息失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": pageResult,
	})
}

// Update 更新物流信息
func (c *LogisticsController) Update(ctx *gin.Context) {
	var logistics model.Logistics
	if err := ctx.ShouldBindJSON(&logistics); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误: " + err.Error(),
		})
		return
	}

	// 验证ID有效性
	if logistics.ID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "物流ID不能为空",
		})
		return
	}

	err := c.service.Update(&logistics)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "更新物流信息失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

// List 查询所有物流信息
func (c *LogisticsController) List(ctx *gin.Context) {
	logisticsList, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "查询物流信息失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": logisticsList,
	})
}

// ConfirmReceipt 确认收货
func (c *LogisticsController) ConfirmReceipt(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的ID",
		})
		return
	}

	// 获取当前物流信息
	logistics, err := c.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "获取物流信息失败: " + err.Error(),
		})
		return
	}

	if logistics == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "未找到对应的物流信息",
		})
		return
	}

	// 设置收货时间为当前时间
	now := time.Now()
	logistics.EndTime = &now

	// 更新物流信息
	err = c.service.Update(logistics)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "确认收货失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "确认收货成功",
	})
}
