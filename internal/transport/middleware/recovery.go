package middleware

import (
	"runtime/debug"
	"time"

	"big_mall_api/internal/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		defer func() {
			if err := recover(); err != nil {
				// 获取请求基本信息
				latency := time.Since(start)
				clientIP := c.ClientIP()
				method := c.Request.Method
				path := c.Request.URL.Path
				statusCode := c.Writer.Status()

				// 获取完整的调用堆栈
				stack := string(debug.Stack())

				// 构建结构化日志字段
				logFields := logrus.Fields{
					"time":       start.Format(time.RFC3339),
					"status":     statusCode,
					"latency":    latency.String(), // 处理耗时
					"client_ip":  clientIP,
					"method":     method,
					"path":       path,
					"error":      err,
					"stack":      stack,
					"query":      c.Request.URL.RawQuery,
					"user_agent": c.Request.UserAgent(),
				}

				// 如果有request_id也记录
				if requestID := c.GetString("X-Request-ID"); requestID != "" {
					logFields["request_id"] = requestID
				}

				// 使用logrus记录错误（Error级别）
				logger.WithFields(logFields).Error("Recovered from panic")

				// 返回标准化错误响应
				c.JSON(500, gin.H{
					"code":      500,
					"message":   "Internal server error",
					"timestamp": start.Unix(),
					//"path":      path,
				})

				c.Abort()
			}
		}()

		c.Next()
	}
}
