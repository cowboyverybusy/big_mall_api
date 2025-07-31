package logic

import (
	"big_mall_api/internal/service"
	"big_mall_api/pkg/mysql"
	"big_mall_api/pkg/redis"
)

type ServerLogic struct {
	storage *service.StorageManager
	//mallServer *service.MallServer
	mainMdb *mysql.Client
	mainRdb *redis.Client
}

func NewServerLogic(storage *service.StorageManager) *ServerLogic {
	mainMbb, _ := storage.GetMySQLClient("main")
	mainRdb, _ := storage.GetRedisClient("main")
	return &ServerLogic{
		storage: storage,
		mainMdb: mainMbb,
		mainRdb: mainRdb,
	}
}
