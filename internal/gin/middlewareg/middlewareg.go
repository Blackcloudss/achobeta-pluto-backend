package middlewareg

import (
	"github.com/gin-gonic/gin"
	"tgwp/log/zlog"
	"tgwp/util/snowflake"
)

const DEFAULTNODE = 1

// AddTraceId 是一个用于生成或获取 Trace ID 的中间件
// 它会将 trace ID 添加到请求的上下文中，并在日志中记录。
func AddTraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的 Trace ID
		traceID := c.GetHeader("X-Request-ID")
		if traceID == "" {

			// 如果没有 Trace ID，生成一个新的,  默认节点为 1
			node, _ := snowflake.NewNode(DEFAULTNODE)
			id := node.Generate()
			traceID = id.String()
		}

		// 将 Trace ID 存入上下文中，方便后续处理使用
		c.Set("traceId", traceID)

		// 在日志中记录 Trace ID
		zlog.CtxInfof(c, "TraceID: %s", traceID)

		// 继续执行下一个中间件或请求处理
		c.Next()
	}
}

// 验证用户是否登录
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 用 header 的 token 进行认证
		token := c.GetHeader("Authorization")
		if token == "" {
			return
		}
		// 验证 token 是否有效
		if !ValidateToken(token) {
			return
		}
		c.Next() // 继续处理请求
	}
}

// 权限校验中间件：检查用户是否有权限访问某个资源
func PermissionMiddleware(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 假设我们有一个函数 GetUserPermissions 检查用户权限
		userPermissions := GetUserPermissions(c)

		if userPermissions == nil {
		}

		c.Next() // 继续处理请求
	}
}

// 错误处理中间件：统一处理异常
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(500, gin.H{"message": "Internal Server Error", "error": err})
			}
		}()
		c.Next() // 继续处理请求
	}
}

// ValidateToken 来验证 token 是否有效
func ValidateToken(token string) bool {
	//  token 校验逻辑
	return token == "valid-token"
}

// GetUserPermissions 获取用户权限组
func GetUserPermissions(c *gin.Context) []string {
	// 权限获取逻辑
	return []string{"view_profile", "edit_profile"}
}
