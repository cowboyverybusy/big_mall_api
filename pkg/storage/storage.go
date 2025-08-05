package storage

import (
	"big_mall_api/configs"
	"big_mall_api/pkg/storage/mysql"
	"big_mall_api/pkg/storage/mysql/container"
	"big_mall_api/pkg/storage/redis"
	"fmt"
	"gorm.io/gorm"
	"sync"
)

type DbManager struct {
	redisMgr *sync.Map
	//mysqlMgr sync.Map
	container *container.RepositoryContainer //存储多个数据库的model实例
}

func NewStorageManager(cfg *configs.Config, models []container.Model) (*DbManager, error) {
	// 初始化所有数据库连接
	dbMap := make(map[string]*gorm.DB)
	for name, dbCfg := range cfg.MySQL {
		db, err := mysql.NewClient(&dbCfg)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to MySQL [%s]: %v", name, err)
		}
		dbMap[name] = db.DB
	}
	fmt.Printf("models:%+v\n", models)
	// 创建容器
	repoContainer := container.NewRepositoryContainer(dbMap)
	// 注册所有模型
	for _, model := range models {
		repoContainer.Register(model)
	}

	// 初始化Redis
	rMgr := &sync.Map{}
	// 初始化 Redis 客户端
	for name, config := range cfg.Redis {
		client := redis.NewClient(&config)
		rMgr.Store(name, client)
	}

	return &DbManager{
		container: repoContainer,
		redisMgr:  rMgr,
	}, nil
}

// GetMySQLClient 根据名字获取 MySQL 客户端
//func (sm *DbManager) GetMySQLClient(name string) (*mysql.Client, bool) {
//	val, ok := sm.mysqlMgr.Load(name)
//	if !ok {
//		return nil, false
//	}
//	client, ok := val.(*mysql.Client)
//	return client, ok
//}
//
//// AddOrUpdateMySQLClient 添加或替换一个 MySQL 客户端
//func (sm *DbManager) AddOrUpdateMySQLClient(name string, client *mysql.Client) {
//	sm.mysqlMgr.Store(name, client)
//}

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

// GetRepository 获取模型仓库
func (sm *DbManager) GetRepository(model container.Model) (interface{}, bool) {
	return container.GetRepository(sm.container, model)
}

// GetDB 获取数据库连接
func (sm *DbManager) GetDB(dbName string) (*gorm.DB, bool) {
	return sm.container.GetDB(dbName)
}
