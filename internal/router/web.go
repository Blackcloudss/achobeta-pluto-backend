package routerg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tgwp/configs"
	"tgwp/global"
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

	// 注册全局中间件（例如获取 Trace ID）
	manager.RequestGlobalMiddleware(r)

	// 创建 RouteManager 实例
	routeManager := manager.NewRouteManager(r)

	// 注册各业务路由组的具体路由
	registerRoutes(routeManager)
	messageRoutes(routeManager)

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
			if token, exists := c.Get(global.AUTH_ENUMS_ATOKEN); exists {
				response.NewResponse(c).Success(token)
			}
			//告诉后面的人如何拿到token里面的数据
			if data, exists := c.Get(global.TOKEN_USER_ID); exists {
				response.NewResponse(c).Success(data)
			}
		})
	})

	// 个人信息相关路由
	routeManager.RegisterProfileRoutes(func(rg *gin.RouterGroup) {

		// example
		rg.GET("/info", func(c *gin.Context) {

		})
		rg.GET("test", api.Test)
	})

	// 团队信息相关路由
	routeManager.RegisterTeamRoutes(func(rg *gin.RouterGroup) {

		//解析 jwt，获取 user_id
		rg.Use(middleware.ReflashAtoken())

		//获得权限组
		rg.GET("/power", api.GetPower)

		// 团队成员管理子路由
		memberGroup := rg.Group("/structure")
		{
			//检验权限
			memberGroup.Use(middleware.PermissionMiddleware())
			// 获取 完整团队架构
			memberGroup.GET("/collection", api.GetTeamStructure)
			//保存 更改了的节点信息
			memberGroup.PUT("/change", api.PutTeamNode)
			//新增团队
			memberGroup.POST("/add", api.PostTeam)
		}
	})
}

// messageRoutes 注册消息相关路由的具体处理函数
func messageRoutes(routeManager *manager.RouteManager) {
	routeManager.HandleMessageRoutes(func(rg *gin.RouterGroup) {
		rg.POST("/send", api.SendMessage)
	})
}
