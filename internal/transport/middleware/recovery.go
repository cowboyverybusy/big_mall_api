package middleware

import (
	"big_mall_api/internal/utils"
	"big_mall_api/internal/utils/logger"
	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Panic recovered", "error", err)
				utils.ErrorResponse(c, 500, "Internal server error", nil)
				c.Abort()
			}
		}()
		c.Next()
	}
}
