package mysql

import (
	"big_mall_api/configs"
	_ "big_mall_api/configs"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Client struct {
	DB *gorm.DB
}

func NewClient(cfg *configs.MySQLConfig) (*Client, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// 获取底层 sql.DB 对象
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	//根据MySQL的wait_timeout属性设置合理的连接池，才不会出现隔一段时间没请求就出现MySQL自动给断开的问题
	//查看全局 wait_timeout：SHOW GLOBAL VARIABLES LIKE 'wait_timeout';
	// 针对MySQL wait_timeout为120秒超时的配置
	sqlDB.SetMaxIdleConns(5)                   // 减少空闲连接数，根据并发量调整
	sqlDB.SetMaxOpenConns(50)                  // 适中的最大连接数，根据并发量调整
	sqlDB.SetConnMaxLifetime(90 * time.Second) // 90秒，ConnMaxLifetime 必须 小于 MySQL 的 wait_timeout
	sqlDB.SetConnMaxIdleTime(60 * time.Second) // 60秒空闲时间 最大空闲时间比MySQL的wait_timeout短10-15%

	//可以考虑通过配置更加灵活设置连接池
	// 从配置读取wait_timeout或使用默认值
	//waitTimeout := dbCfg.WaitTimeout
	//if waitTimeout == 0 {
	//	waitTimeout = 280 // 默认比MySQL常规配置300秒短
	//}
	//// 动态设置连接生命周期
	//sqlDB.SetConnMaxLifetime(time.Duration(waitTimeout) * time.Second)
	//sqlDB.SetConnMaxIdleTime(time.Duration(waitTimeout/2) * time.Second)
	//sqlDB.SetMaxIdleConns(dbCfg.MaxIdleConns)
	//sqlDB.SetMaxOpenConns(dbCfg.MaxOpenConns)

	return &Client{DB: db}, nil
}

func (c *Client) Close() error {
	sqlDB, err := c.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
