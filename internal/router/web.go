package routerg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tgwp/configs"
	"tgwp/internal/api"
	"tgwp/internal/manager"
	"tgwp/internal/middleware"
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
		//发送验证码
		rg.POST("/code", api.GetCode)
		//用验证码登录
		rg.POST("/login", api.LoginWithCode)
		//是否可以自动登录
		rg.POST("/auto", api.CheckAutoLogin)
		//用户自行退出登录
		rg.DELETE("/exit", api.ExitSystem)
	})

	// 展示常用设备页面相关操作路由
	routeManager.RegisterDevicesRoutes(func(rg *gin.RouterGroup) {
		//使用token验证中间件
		rg.Use(middleware.ReflashAtoken())
		//移除常用设备
		rg.DELETE("/remove", api.RemoveDevice)
		//展示常用设备
		rg.GET("/show", api.ShowDevices)
		//修改设备名称
		rg.PUT("/modify", api.ModifyDeviceName)
	})
	// 个人信息相关路由
	routeManager.RegisterProfileRoutes(func(rg *gin.RouterGroup) {
		//测试
		rg.GET("test", api.Test)
	})

	// 团队信息相关路由
	routeManager.RegisterTeamRoutes(func(rg *gin.RouterGroup) {
		//使用token验证中间件  正式使用时需取消注释
		//rg.Use(middleware.ReflashAtoken())
		//获得权限组
		rg.GET("/power", api.GetPower)
		// 团队架构管理子路由
		TeamStructure := rg.Group("/structure")
		{
			//检验权限
			//TeamStructure.Use(middleware.PermissionMiddleware())
			// 获取 该团队架构全部节点
			TeamStructure.GET("/collection", api.GetTeamStructure)
			//保存 更改了的节点信息
			TeamStructure.PUT("/change", api.PutTeamNode)
			//新增团队
			TeamStructure.POST("/create", api.CreateTeam)
		}

		// 团队成员列表子路由
		MemberList := rg.Group("/memberlist")
		{
			//查询团队内成员列表--成员简单信息
			MemberList.GET("/get", api.GetTeamMemberlist)
			//使用权限校验中间件
			//MemberList.Use(middleware.PermissionMiddleware())
			//新增成员
			MemberList.POST("/create", api.CreateTeamMember)
			//删除成员                     // user_id 之后要删除
			MemberList.DELETE("/delete/:team_id/:member_id", api.DeleteTeamMember)
		}

		// 团队成员信息管理子路由
		MemberMsg := rg.Group("/membermsg")
		{
			//查询成员详细信息
			MemberMsg.GET("/details", api.GetMemberDetail)
			//给成员点赞/取消赞
			MemberMsg.PUT("/like", api.PutLikeCount)
			//使用权限校验中间件
			MemberMsg.Use(middleware.PermissionMiddleware())
			//编辑成员信息
			MemberMsg.PUT("/change", api.PutTeamMember)
		}
	})

	// 消息相关路由组
	routeManager.RegisterMessageRoutes(func(rg *gin.RouterGroup) {
		// 建立消息源
		rg.POST("/set", api.SetMessage)
		// 标记已读
		rg.POST("/markread", api.MarkReadMessage)
		//使用token验证中间件
		rg.Use(middleware.ReflashAtoken())
		// 将消息与用户绑定
		rg.POST("/join", api.JoinMessage)
		// 获取消息列表(分页)
		rg.GET("/get", api.GetMessage)
		// 一键发送消息
		rg.POST("/send", api.SendMessage)
		// 标记全部消息已读
		rg.POST("/markread-all", api.MarkReadAllMessage)
	})

	// 飞书相关路由组
	routeManager.RegisterFeiShuRoutes(func(rg *gin.RouterGroup) {
		//使用token验证中间件
		rg.Use(middleware.ReflashAtoken())
		// 获取飞书二维表格
		rg.GET("/get", api.GetFeiShuList)
	})

	// 个人中心相关路由组
	routeManager.RegisterUserProfileRoutes(func(rg *gin.RouterGroup) {
		//使用token验证中间件
		rg.Use(middleware.ReflashAtoken())
		//查询成员详细信息
		rg.GET("/details", api.GetUserDetail)
		//给成员点赞/取消赞
		rg.PUT("/like", api.PutUserLikeCount)
	})
}
