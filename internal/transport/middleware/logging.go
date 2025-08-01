package middleware

import (
	"big_mall_api/internal/utils/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("==== come in LoggingMiddleware ====")
		// 请求开始时间（高精度）
		start := time.Now()

		// 获取请求信息
		path := c.Request.URL.Path
		rawQuery := c.Request.URL.RawQuery
		method := c.Request.Method
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		// 处理请求
		c.Next()

		// 计算延迟（毫秒精度）
		latency := time.Since(start)
		latencyMs := float64(latency.Nanoseconds()) / 1e6 // 转换为毫秒

		// 获取响应信息
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
		respSize := c.Writer.Size() // 响应体大小

		// 构建完整路径
		fullPath := path
		if rawQuery != "" {
			fullPath = path + "?" + rawQuery
		}

		// 结构化日志输出
		logger.WithFields(logrus.Fields{
			"method":     method,
			"path":       fullPath,
			"status":     statusCode,
			"latency":    fmt.Sprintf("%.3fms", latencyMs), // 保留3位小数
			"client_ip":  clientIP,
			"user_agent": userAgent,
			"error":      errorMessage,
			"resp_size":  respSize, // 响应大小(字节)
			//"timestamp":  time.Now().Format(time.RFC3339),
		}).Info("HTTP Request")
	}
}
