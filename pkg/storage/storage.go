package storage

import (
	"big_mall_api/configs"
	"big_mall_api/pkg/storage/mysql"
	"big_mall_api/pkg/storage/redis"
	"fmt"
	"sync"
)

type DbManager struct {
	redisMgr sync.Map
	mysqlMgr sync.Map
}

func NewStorageManager(cfg *configs.Config) (*DbManager, error) {
	storage := &DbManager{}
	// 初始化 MySQL 客户端
	for name, config := range cfg.MySQL {
		client, err := mysql.NewClient(&config)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to MySQL [%s]: %v", name, err)
		}
		storage.mysqlMgr.Store(name, client)
	}

	// 初始化 Redis 客户端
	for name, config := range cfg.Redis {
		client := redis.NewClient(&config)
		storage.redisMgr.Store(name, client)
	}

	return storage, nil
}

// GetMySQLClient 根据名字获取 MySQL 客户端
func (sm *DbManager) GetMySQLClient(name string) (*mysql.Client, bool) {
	val, ok := sm.mysqlMgr.Load(name)
	if !ok {
		return nil, false
	}
	client, ok := val.(*mysql.Client)
	return client, ok
}

// AddOrUpdateMySQLClient 添加或替换一个 MySQL 客户端
func (sm *DbManager) AddOrUpdateMySQLClient(name string, client *mysql.Client) {
	sm.mysqlMgr.Store(name, client)
}

// GetRedisClient 根据名字获取 Redis 客户端
func (sm *DbManager) GetRedisClient(name string) (*redis.Client, bool) {
	val, ok := sm.redisMgr.Load(name)
	if !ok {
		return nil, false
	}
	client, ok := val.(*redis.Client)
	return client, ok
}

// AddOrUpdateRedisClient 添加或替换一个 Redis 客户端
func (sm *DbManager) AddOrUpdateRedisClient(name string, client *redis.Client) {
	sm.redisMgr.Store(name, client)
}
