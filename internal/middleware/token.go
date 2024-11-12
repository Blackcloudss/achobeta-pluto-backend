package middleware

import (
	"github.com/gin-gonic/gin"
	"tgwp/global"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/log/zlog"
	"tgwp/util"
)

// 用于刷新token
func ReflashAtoken() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := zlog.GetCtxFromGin(c)
		token := c.GetHeader("Authorization")
		if token == "" {
			zlog.CtxErrorf(ctx, `token is empty`)
			c.Abort()
			return
		}
		//解析token是否有效，并取出上一次的值
		data, err := util.IdentifyToken(ctx, token)
		if err != nil {
			zlog.CtxErrorf(ctx, "ReflashAtoken err:%v", err)
			response.NewResponse(c).Error(response.TOKEN_IS_EXPIRED)
			//对应token无效，直接让他返回
			c.Abort()
			return
		}
		//判断其是否为atoken
		if data.Class != global.AUTH_ENUMS_ATOKEN {
			response.NewResponse(c).Error(response.PARAM_TYPE_ERROR)
			c.Abort()
			return
		}
		//生成新的token
		resp, err := logic.NewTokenLogic().AtokenLogic(ctx, data)
		if err != nil {
			zlog.CtxErrorf(ctx, "ReflashAtoken err:%v", err)
			c.Abort()
			return
		}
		//将值传递给后面用
		c.Set("Token", resp.Atoken)
		c.Next()
	}
}
