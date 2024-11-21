package middleware

import (
	"github.com/gin-gonic/gin"
	"tgwp/global"
	"tgwp/internal/logic"
	"tgwp/internal/repo"
	"tgwp/internal/response"
	"tgwp/log/zlog"
	"tgwp/util"
)

// ReflashAtoken
//
//	@Description: 用于刷新token
//	@return gin.HandlerFunc
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
		//判断token内部的签名是否为空，只有点了自动登录才不为空
		//点了自动登录,得去签名表判断签名是否有效
		if data.Issuer != "" {
			//进行判断
			err = repo.NewSignRepo(global.DB).CompareSign(data.Issuer)
			if err != nil {
				//表明找不到issuer相等的，即atoken是无效的
				zlog.CtxErrorf(ctx, "ReflashAtoken err:%v", err)
				response.NewResponse(c).Error(response.PARAM_NOT_VALID)
				c.Abort()
				return
			}
		}
		//将token内部数据传下去
		c.Set(global.TOKEN_USER_ID, data.Userid)
		//生成新的token
		resp, err := logic.NewTokenLogic().GenAtoken(ctx, data)
		if err != nil {
			zlog.CtxErrorf(ctx, "ReflashAtoken err:%v", err)
			c.Abort()
			return
		}
		//将值传递给后面用
		c.Set(global.AUTH_ENUMS_ATOKEN, resp.Atoken)
		c.Next()
	}
}
