package logic

import (
	"big_mall_api/pkg/storage"
	"big_mall_api/pkg/storage/mysql"
	"big_mall_api/pkg/storage/redis"
)

type ServerLogic struct {
	dbMgr   *storage.DbManager
	mainMdb *mysql.Client
	mainRdb *redis.Client
	//后续可以添加多个mysql、redis
}

func NewServerLogic(dbMgr *storage.DbManager) *ServerLogic {
	mainMbb, _ := dbMgr.GetMySQLClient("main")
	mainRdb, _ := dbMgr.GetRedisClient("main")
	return &ServerLogic{
		dbMgr:   dbMgr,
		mainMdb: mainMbb,
		mainRdb: mainRdb,
	}
}
