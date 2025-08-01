package service

import (
	"big_mall_api/internal/transport/middleware"
	"github.com/gin-gonic/gin"
)

func (s *MallServer) setupApiRoutes() {
	s.SetupAuthAPIRoutes()
	s.SetupPublicAPIRoutes()
}

// SetupAuthAPIRoutes 设置需要认证的API路由
func (s *MallServer) SetupAuthAPIRoutes() {
	// 创建API组并应用认证中间件
	authApi := s.engine.Group("/api")

	authApi.Use(middleware.AuthMiddleware())    // 先应用授权中间件
	authApi.Use(middleware.LoggingMiddleware()) //认证通过后再记录日志

	// 用户相关路由
	userGroup := authApi.Group("/user")
	{
		userGroup.GET("/list", s.handlers.ListUsers)
		userGroup.GET("/getUserById", s.handlers.GetUser)
		userGroup.POST("/create", s.handlers.CreateUser) // 通常创建用户用POST方法
	}

	// 可以继续添加其他需要认证的路由组
	// orderGroup := authApi.Group("/order")
	// {
	//     orderGroup.GET("/list", s.handlers.ListOrders)
	//     ...
	// }
}

// SetupPublicAPIRoutes 设置公开访问的API路由
func (s *MallServer) SetupPublicAPIRoutes() {

	s.engine.Use(middleware.LoggingMiddleware()) //日志中间件

	s.engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	//s.engine.GET("/health", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"status": "healthy"})
	//})
	//
	//// 登录路由（通常不需要认证）
	//s.engine.POST("/login", s.handlers.Login)
	//
	//// 注册路由（通常不需要认证）
	//s.engine.POST("/register", s.handlers.Register)
	//
	//// 公开产品信息
	//productGroup := s.engine.Group("/products")
	//{
	//	productGroup.GET("/list", s.handlers.ListProducts)
	//	productGroup.GET("/detail", s.handlers.GetProductDetail)
	//}
}

func (s *MallServer) setupMiddleware() {
	//Prometheus 监控应放在最外层 需要统计完整请求链路的耗时（包括其他中间件的执行时间）如果放在内层，会漏测外层中间件的性能指标
	s.engine.Use(middleware.PrometheusMiddleware())
	s.engine.Use(middleware.CORSMiddleware())
	s.engine.Use(middleware.RecoveryMiddleware()) // Panic恢复应在最靠近业务的地方
}
