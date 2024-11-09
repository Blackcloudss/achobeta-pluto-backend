package manager

import (
	"github.com/gin-gonic/gin"
)

//主要管理路由组和中间件的注册

// PathHandler 是一个用于注册路由组的函数类型
type PathHandler func(rg *gin.RouterGroup)

// Middleware 是一个用于生成中间件的函数类型
type Middleware func() gin.HandlerFunc

// RouteManager 管理不同的路由组，按业务功能分组
type RouteManager struct {
	LoginRoutes   *gin.RouterGroup // 登录相关的路由组
	ProfileRoutes *gin.RouterGroup // 个人信息相关的路由组
	TeamRoutes    *gin.RouterGroup // 团队信息相关的路由组
}

// NewRouteManager 创建一个新的 RouteManager 实例，包含各业务功能的路由组
func NewRouteManager(router *gin.Engine) *RouteManager {
	return &RouteManager{
		LoginRoutes:   router.Group("/api/login"),   // 初始化登录路由组
		ProfileRoutes: router.Group("/api/profile"), // 初始化个人信息路由组
		TeamRoutes:    router.Group("/api/team"),    // 初始化团队信息路由组
	}
}

// RegisterLoginRoutes 注册登录相关的路由处理函数
func (rm *RouteManager) RegisterLoginRoutes(handler PathHandler) {
	handler(rm.LoginRoutes)
}

// RegisterProfileRoutes 注册个人信息相关的路由处理函数
func (rm *RouteManager) RegisterProfileRoutes(handler PathHandler) {
	handler(rm.ProfileRoutes)
}

// RegisterTeamRoutes 注册团队信息相关的路由处理函数
func (rm *RouteManager) RegisterTeamRoutes(handler PathHandler) {
	handler(rm.TeamRoutes)
}

// RegisterMiddleware 根据组名为对应的路由组注册中间件
// group 参数为 "login"、"profile" 或 "team"，分别对应不同的路由组
func (rm *RouteManager) RegisterMiddleware(group string, middleware Middleware) {
	switch group {
	case "login":
		rm.LoginRoutes.Use(middleware())
	case "profile":
		rm.ProfileRoutes.Use(middleware())
	case "team":
		rm.TeamRoutes.Use(middleware())
	}
}

// RequestGlobalMiddleware 注册全局中间件，应用于所有路由
func RequestGlobalMiddleware(r *gin.Engine) {

}
