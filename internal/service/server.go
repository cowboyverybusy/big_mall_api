package service

import (
	"big_mall_api/configs"
	"big_mall_api/internal/logic"
	"big_mall_api/internal/transport/handler"
	"big_mall_api/internal/transport/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MallServer struct {
	engine         *gin.Engine
	storageManager *StorageManager
	cfg            *configs.Config
	logger         *logrus.Logger
	handlers       *handler.MallServerHandler
	logics         *logic.ServerLogic
}

func NewMallServer(cfg *configs.Config, storageMgr *StorageManager, logics *logic.ServerLogic, logger *logrus.Logger) *MallServer {
	gin.SetMode(cfg.Server.Mode)
	engine := gin.New()

	return &MallServer{
		engine:         engine,
		storageManager: storageMgr,
		cfg:            cfg,
		logger:         logger,
		logics:         logics,
	}
}

func (s *MallServer) setupMiddleware() {
	s.engine.Use(middleware.RecoveryMiddleware())
	s.engine.Use(middleware.LoggingMiddleware())
	s.engine.Use(middleware.CORSMiddleware())
	s.engine.Use(middleware.PrometheusMiddleware())
}

func (s *MallServer) setupRoutes() {
	s.SetupAPIRoutes()
	// 初始化Repository层
	//mysqlRepo := repository.NewMySQLRepository(mysqlClient)
	//redisRepo := repository.NewRedisRepository(redisClient)
	//esRepo := repository.NewESRepository(esClient)

	// 初始化Service层
	//userService := service.NewUserService(mysqlRepo, redisRepo, esRepo)

	// 初始化Controller层
	//userController := controller.NewUserController(userService)

	// 初始化Handler层
	//userHandler := handler.NewUserHandler(userController)

	// 健康检查
	//s.engine.GET("/health", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"status": "ok"})
	//})
	//
	//// Prometheus metrics
	//s.engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	//
	//// API路由组
	//v1 := s.engine.Group("/api/v1")
	//
	//// 公开路由
	//public := v1.Group("/")
	//userHandler.RegisterRoutes(public)
	//
	//// 需要认证的路由
	//protected := v1.Group("/")
	//protected.Use(middleware.AuthMiddleware())
	// 在这里添加需要认证的路由
}

func (s *MallServer) registerHandle() {
	s.handlers = handler.NewMallServerHandler(s.logics)
}

func (s *MallServer) Run() error {
	// 设置中间件
	s.setupMiddleware()

	// 设置handler
	s.registerHandle()

	// 设置路由
	s.setupRoutes()

	// 启动服务器
	return s.engine.Run(":" + s.cfg.Server.Port)
}
