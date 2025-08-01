package mysql

import (
	"big_mall_api/configs"
	_ "big_mall_api/configs"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	return &Client{DB: db}, nil
}

func (c *Client) Close() error {
	sqlDB, err := c.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
