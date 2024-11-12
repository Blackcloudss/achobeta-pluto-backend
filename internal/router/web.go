package routerg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tgwp/configs"
	"tgwp/internal/api"
	"tgwp/internal/manager"
	"tgwp/internal/middleware"
	"tgwp/internal/response"
	"tgwp/log/zlog"
)

// RunServer 启动服务器 路由层
func RunServer() {
	r, err := listen()
	if err != nil {
		zlog.Errorf("Listen error: %v", err)
		panic(err.Error())
	}
	r.Run(fmt.Sprintf("%s:%d", configs.Conf.App.Host, configs.Conf.App.Port)) // 启动 Gin 服务器
}

// listen 配置 Gin 服务器
func listen() (*gin.Engine, error) {
	r := gin.Default() // 创建默认的 Gin 引擎

	// 注册全局中间件（录入例如或获取 Trace ID）
	manager.RequestGlobalMiddleware(r)

	// 创建 RouteManager 实例
	routeManager := manager.NewRouteManager(r)

	// 注册各业务路由组的具体路由
	registerRoutes(routeManager)

	return r, nil
}

// registerRoutes 注册各业务路由的具体处理函数
func registerRoutes(routeManager *manager.RouteManager) {
	//通用功能相关路由
	routeManager.RegisterCommonRoutes(func(rg *gin.RouterGroup) {
		rg.POST("/rtoken", api.ReflashRtoken)
	})
	// 登录相关路由
	routeManager.RegisterLoginRoutes(func(rg *gin.RouterGroup) {
		rg.POST("/login", api.LoginWithCode)
		rg.POST("/code", api.GetCode)
		rg.GET("/test", middleware.ReflashAtoken(), func(c *gin.Context) {
			if token, exists := c.Get("Token"); exists {
				response.NewResponse(c).Success(token)
			}
		})
	})

	// 个人信息相关路由
	routeManager.RegisterProfileRoutes(func(rg *gin.RouterGroup) {
		// example
		rg.Use(middleware.AuthMiddleware()) // 认证中间件

		// example
		rg.GET("/info", func(c *gin.Context) {

		})
		rg.GET("test", api.Test)
	})

	// 团队信息相关路由
	routeManager.RegisterTeamRoutes(func(rg *gin.RouterGroup) {
		// example:认证和权限校验中间件
		rg.Use(middleware.AuthMiddleware())                  // 认证中间件
		rg.Use(middleware.PermissionMiddleware("view_team")) // 权限校验中间件
		rg.GET("/list", func(c *gin.Context) {

		})

		// 团队成员管理子路由
		memberGroup := rg.Group("/members")
		memberGroup.Use() // 注册成员管理中间件
		{
			memberGroup.GET("/list", func(c *gin.Context) {

			})
		}
	})

}
