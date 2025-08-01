package middleware

import (
	"big_mall_api/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		fmt.Printf("======= token:%+v\n", token)
		if token == "" {
			utils.ErrorResponse(c, 401, "Authorization header required", nil)
			c.Abort()
			return
		}

		// 简单的token验证示例
		if !strings.HasPrefix(token, "Bearer ") {
			utils.ErrorResponse(c, 401, "Invalid authorization format", nil)
			c.Abort()
			return
		}

		// TODO: 实际的token验证逻辑
		// 这里可以验证JWT token，查询数据库等

		c.Next()
	}
}
