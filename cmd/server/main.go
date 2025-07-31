package main

import (
	"big_mall_api/configs"
	"big_mall_api/internal/logic"
	"big_mall_api/internal/service"
	"big_mall_api/internal/utils/logger"
	"log"
)

func main() {
	// 加载配置
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化日志
	globalLogger, err := logger.Init(&cfg.Log)
	if err != nil {
		log.Fatalf("Failed to init logger: %v", err)
	}

	// 初始化存储系统（MySQL、redis）
	storageMgr, err := service.NewStorageManager(cfg)
	if err != nil {
		log.Fatalf("Failed to init storageMannager: %v", err)
	}

	serverLogic := logic.NewServerLogic(storageMgr)
	// 启动服务器
	server := service.NewMallServer(cfg, storageMgr, serverLogic, globalLogger)
	if err := server.Run(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
