package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"tgwp/global"
	"tgwp/internal/response"
	"tgwp/log/zlog"
	"tgwp/util"
	"time"
)

// 用作验证码处理
func PostCode(ctx context.Context, phone string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//生成6位数字的验证码
		code := util.RandomCode()
		text := fmt.Sprintf("你的验证码是%s", code)
		//将验证码放入redis5分钟
		global.Rdb.Set(ctx, fmt.Sprintf(global.REDIS_PHONE_CODE, phone), code, time.Second*300)
		//发送验证码
		if err := PostMessage(text, phone); err != nil {
			zlog.CtxErrorf(ctx, "PostMessage err: %v", err)
		}
		response.NewResponse(c).Success(code)
		//防刷处理
		//数字3用作占位参数，后续可以对其进行扩充业务，如限制其重新获取验证码次数
		global.Rdb.Set(ctx, fmt.Sprintf(global.REDIS_PHONE, phone), 3, time.Second*60)
	}
}
func PostMessage(text string, phone string) error {
	return nil
}
