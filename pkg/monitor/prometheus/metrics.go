package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HTTPRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "big_mall",
			Subsystem: "http",
			Name:      "http_requests_total",
			Help:      "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status_code"},
	)

	HTTPRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "big_mall",
			Subsystem: "http",
			Name:      "http_request_duration_seconds",
			Help:      "HTTP request duration in seconds",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	//OnlineUsers = promauto.NewGauge(
	//	prometheus.GaugeOpts{
	//		Namespace: "big_mall",
	//		Subsystem: "business",
	//		Name:      "online_users",
	//		Help:      "Number of currently online users",
	//	},
	//)
	//
	//Orders = promauto.NewCounterVec(
	//	prometheus.CounterOpts{
	//		Namespace: "big_mall",
	//		Subsystem: "business",
	//		Name:      "orders_total",
	//		Help:      "Total number of orders",
	//	},
	//	[]string{"status"},
	//)

	//OrderValue = promauto.NewHistogramVec(
	//	prometheus.HistogramOpts{
	//		Namespace: "big_mall",
	//		Subsystem: "business",
	//		Name:      "order_value",
	//		Help:      "Order value distribution",
	//		Buckets:   []float64{10, 50, 100, 250, 500, 1000, 2500, 5000, 10000},
	//	},
	//	[]string{"category"},
	//)

	// 系统资源指标
	GoroutinesCount = promauto.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "big_mall",
			Subsystem: "system",
			Name:      "goroutines",
			Help:      "Number of goroutines",
		},
	)

	MemoryUsage = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "big_mall",
			Subsystem: "system",
			Name:      "memory_usage_bytes",
			Help:      "Memory usage in bytes",
		},
		[]string{"type"},
	)

	CPUUsage = promauto.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "big_mall",
			Subsystem: "system",
			Name:      "cpu_usage_percent",
			Help:      "CPU usage percentage",
		},
	)
)
