package service

import (
	"big_mall_api/internal/transport/middleware"
	"github.com/gin-gonic/gin"
)

func (s *MallServer) SetupAPIRoutes() {
	// 需要认证的路由
	authApi := s.engine.Group("/api")
	{
		authApi.POST("/user/list", s.handlers.ListUsers)
		authApi.POST("/user/getUserById", s.handlers.GetUser)
		authApi.POST("/user/createUser", s.handlers.CreateUser)
	}
	//使用认证中间件
	authApi.Use(middleware.AuthMiddleware())

	//公开路由
	s.engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}
