package monitor

import (
	"context"
	"log"
	"runtime"
	"time"
)

type SystemMetricCollector struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewMetricsSystemCollector() *SystemMetricCollector {
	ctx, cancel := context.WithCancel(context.Background())
	return &SystemMetricCollector{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (sc *SystemMetricCollector) Start() {
	// 首次立即采集
	sc.collectMetrics()
	ticker := time.NewTicker(15 * time.Second)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-sc.ctx.Done():
				return
			case <-ticker.C:
				log.Println("采集一次数据")
				sc.collectMetrics()
			}
		}
	}()
}

func (sc *SystemMetricCollector) Stop() {
	log.Println("退出监控")
	sc.cancel()
}

func (sc *SystemMetricCollector) collectMetrics() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// Goroutines数量
	GoroutinesCount.Set(float64(runtime.NumGoroutine()))

	// 2. 采集内存指标
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	MemoryUsage.WithLabelValues("alloc").Set(float64(memStats.Alloc))
	MemoryUsage.WithLabelValues("total_alloc").Set(float64(memStats.TotalAlloc))
	MemoryUsage.WithLabelValues("sys").Set(float64(memStats.Sys))
	MemoryUsage.WithLabelValues("heap_alloc").Set(float64(memStats.HeapAlloc))
	MemoryUsage.WithLabelValues("heap_sys").Set(float64(memStats.HeapSys))
	MemoryUsage.WithLabelValues("stack_sys").Set(float64(memStats.StackSys))
	MemoryUsage.WithLabelValues("gc_sys").Set(float64(memStats.GCSys))
}
