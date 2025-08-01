package middleware

import (
	prom "big_mall_api/pkg/monitor/prometheus"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("==== come in PrometheusMiddleware ====")
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		statusCode := strconv.Itoa(c.Writer.Status())

		prom.HTTPRequestsTotal.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
			statusCode,
		).Inc()

		prom.HTTPRequestDuration.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
		).Observe(duration.Seconds())
	}
}
