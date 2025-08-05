package service

import (
	"big_mall_api/configs"
	"big_mall_api/internal/transport/handler"
	monitor "big_mall_api/pkg/monitor/prometheus"
	"big_mall_api/pkg/storage"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
)

type MallServer struct {
	engine              *gin.Engine
	storageManager      *storage.DbManager
	cfg                 *configs.Config
	logger              *logrus.Logger
	handlers            *handler.MallServerHandler
	metricsServer       *http.Server
	sysMetricsCollector *monitor.SystemMetricCollector
}

func NewMallServer(cfg *configs.Config, storageMgr *storage.DbManager, logger *logrus.Logger) *MallServer {
	gin.SetMode(cfg.Server.Mode)
	engine := gin.New()

	return &MallServer{
		engine:              engine,
		storageManager:      storageMgr,
		cfg:                 cfg,
		logger:              logger,
		sysMetricsCollector: monitor.NewMetricsSystemCollector(),
	}
}

func (s *MallServer) registerHandle() {
	s.handlers = handler.NewMallServerHandler(s.storageManager)
}

func (s *MallServer) Run() error {
	// 启动系统指标收集器
	s.sysMetricsCollector.Start()
	defer s.sysMetricsCollector.Stop()

	// 设置handler(在路由设置之前先执行这个)
	s.registerHandle()

	// 启动Prometheus指标服务器
	s.startMetricsServer()

	// 设置中间件(针对所有路由)
	s.setupMiddleware()

	// 设置路由
	s.setupApiRoutes()

	// 启动服务器
	return s.engine.Run(":" + s.cfg.Server.Port)
}

func (s *MallServer) startMetricsServer() {
	if !s.cfg.Prometheus.Enabled {
		return
	}

	mux := http.NewServeMux()
	mux.Handle(s.cfg.Prometheus.MetricPath, promhttp.Handler())

	s.metricsServer = &http.Server{
		//Addr:    ":9090", // Prometheus指标服务器端口
		Addr:    fmt.Sprintf(":%s", s.cfg.Prometheus.Port), // Prometheus指标服务器端口
		Handler: mux,
	}

	go func() {
		s.logger.Infof("Starting Prometheus metrics server on :%s", s.cfg.Prometheus.Port)
		if err := s.metricsServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Errorf("Metrics server error: %v", err)
		}
	}()
}

func (s *MallServer) Shutdown(ctx context.Context) error {
	// 停止系统指标收集器
	s.sysMetricsCollector.Stop()

	// 关闭指标服务器
	if s.metricsServer != nil {
		if err := s.metricsServer.Shutdown(ctx); err != nil {
			s.logger.Errorf("Error shutting down metrics server: %v", err)
		}
	}

	// 关闭主服务器
	return nil
}
