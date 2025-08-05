package logic

import (
	"big_mall_api/pkg/storage"
	"big_mall_api/pkg/storage/redis"
	"gorm.io/gorm"
)

type ServerLogic struct {
	dbMgr   *storage.DbManager
	mainMdb *gorm.DB
	mainRdb *redis.Client
	//后续可以添加多个mysql、redis
}

func NewServerLogic(dbMgr *storage.DbManager) *ServerLogic {
	mainMbb, _ := dbMgr.GetDB("main")
	mainRdb, _ := dbMgr.GetRedisClient("main")
	return &ServerLogic{
		dbMgr:   dbMgr,
		mainMdb: mainMbb,
		mainRdb: mainRdb,
	}
}
