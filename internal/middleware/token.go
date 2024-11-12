package middleware

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/log/zlog"
	"tgwp/util"
)

// 用于刷新token
func ReflashToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		zlog.GetCtxFromGin(c)
		token := c.GetHeader("Authorization")
		if token == "" {
			zlog.CtxErrorf(c, `token is empty`)
			c.Abort()
			return
		}
		//解析token是否有效，并取出上一次的值
		data, err := util.IdentifyToken(c, token)
		if err != nil {
			zlog.CtxErrorf(c, "ReflashToken err:%v", err)
			response.NewResponse(c).Error(response.TOKEN_IS_EXPIRED)
			//对应token无效，直接让他返回
			c.Abort()
			return
		}
		//生成新的token
		resp, err := logic.NewTokenLogic().TokenLogic(c, data)
		if err != nil {
			zlog.CtxErrorf(c, "ReflashToken err:%v", err)
			c.Abort()
			return
		}
		//将值传递给后面用
		c.Set("Token", resp)
		c.Next()
	}
}
