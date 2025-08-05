package main

import (
	"big_mall_api/configs"
	"big_mall_api/internal/model"
	"big_mall_api/internal/service"
	"big_mall_api/internal/utils/logger"
	"big_mall_api/pkg/storage"
	"log"
)

func main() {
	// 加载配置
	cfg, err := configs.LoadConfig("../../configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化日志
	globalLogger, err := logger.Init(&cfg.Log)
	if err != nil {
		log.Fatalf("Failed to init logger: %v", err)
	}

	// 初始化存储系统（MySQL、redis）
	storageMgr, err := storage.NewStorageManager(cfg, model.GetContainerModelList())
	if err != nil {
		log.Fatalf("Failed to init storageMannager: %v", err)
	}

	// 启动服务器
	server := service.NewMallServer(cfg, storageMgr, globalLogger)
	if err := server.Run(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
