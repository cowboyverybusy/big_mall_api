package service

import (
	"big_mall_api/configs"
	"big_mall_api/internal/transport/handler"
	"big_mall_api/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MallServer struct {
	engine         *gin.Engine
	storageManager *storage.DbManager
	cfg            *configs.Config
	logger         *logrus.Logger
	handlers       *handler.MallServerHandler
}

func NewMallServer(cfg *configs.Config, storageMgr *storage.DbManager, logger *logrus.Logger) *MallServer {
	gin.SetMode(cfg.Server.Mode)
	engine := gin.New()

	return &MallServer{
		engine:         engine,
		storageManager: storageMgr,
		cfg:            cfg,
		logger:         logger,
	}
}

func (s *MallServer) registerHandle() {
	s.handlers = handler.NewMallServerHandler(s.storageManager)
}

func (s *MallServer) Run() error {
	// 设置中间件(针对所有路由)
	s.setupMiddleware()

	// 设置路由
	s.setupApiRoutes()

	// 设置handler
	s.registerHandle()

	// 启动服务器
	return s.engine.Run(":" + s.cfg.Server.Port)
}
