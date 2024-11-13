package middlewareh

import (
	"github.com/gin-gonic/gin"
	"tgwp/global"
	"tgwp/log/zlog"
	"tgwp/util/snowflake"
)

// AddTraceId 是一个用于生成或获取 Trace ID 的中间件
// 它会将 trace ID 添加到请求的上下文中，并在日志中记录。
func AddTraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的 Trace ID
		traceID := c.GetHeader("X-Request-ID")
		if traceID == "" {
			//如果没有 Trace ID，生成一个新的,
			node, _ := snowflake.NewNode(global.Default_Node_ID)
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
