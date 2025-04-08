package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"agricultural_product_gin/utils"
)

// JWTMiddleware JWT验证中间件
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "未提供认证令牌",
			})
			c.Abort()
			return
		}

		// 检查Bearer前缀
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "令牌格式错误",
			})
			c.Abort()
			return
		}

		// 解析Token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "无效的令牌",
			})
			c.Abort()
			return
		}

		// 将用户信息保存到ThreadLocal
		threadLocal := utils.GetUserLocal()
		threadLocal.Set("userID", claims.UserID)
		threadLocal.Set("username", claims.Username)

		c.Next()
	}
}
