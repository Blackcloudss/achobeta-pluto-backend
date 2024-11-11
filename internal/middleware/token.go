package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"tgwp/global"
	"tgwp/internal/response"
	"tgwp/log/zlog"
	"tgwp/util"
)

// 用于刷新token
func ReflashToken(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			zlog.CtxErrorf(ctx, `token is empty`)
			return
		}
		//解析token是否有效，并取出上一次的值
		var data util.TokenData
		data, err := util.IdentifyToken(ctx, token)
		if err != nil {
			zlog.CtxErrorf(ctx, "ReflashToken err:%v", err)
			response.NewResponse(c).Error(response.TOKEN_IS_EXPIRED)
			//对应token无效，直接让他返回
			c.Abort()
			return
		}
		//生成新的token
		if data.Class == "atoken" {
			atoken, err := util.GenToken(data)
			if err != nil {
				zlog.CtxErrorf(ctx, "ReflashToken err:%v", err)
			}
			temp := map[string]string{
				"atoken": atoken,
			}
			response.NewResponse(c).Success(temp)
		} else {
			rtoken, err := util.GenToken(data)
			if err != nil {
				zlog.CtxErrorf(ctx, "ReflashToken err:%v", err)
			}
			data.Time = global.ATOKEN_EFFECTIVE_TIME
			data.Class = "atoken"
			atoken, err := util.GenToken(data)
			if err != nil {
				zlog.CtxErrorf(ctx, "ReflashToken err:%v", err)
			}
			temp := map[string]string{
				"atoken": atoken,
				"rtoken": rtoken,
			}
			response.NewResponse(c).Success(temp)
		}
	}
}
