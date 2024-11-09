package handler

import (
	"context"
	"fmt"
	"tgwp/global"
	"tgwp/log/zlog"
	"tgwp/util"
	"time"
)

// 用作验证码处理
func PostCode(ctx context.Context, phone string) error {
	//生成6位数字的验证码
	code := util.RandomCode()
	text := fmt.Sprintf("你的验证码是%s", code)
	//将验证码放入redis5分钟
	err := global.Rdb.Set(ctx, fmt.Sprintf(global.REDIS_PHONE_CODE, phone), code, time.Second*300).Err()
	if err != nil {
		zlog.CtxErrorf(ctx, "Store the verification code err: %v", err)
	}
	//防刷处理
	//数字0用作占位参数
	err = global.Rdb.Set(ctx, fmt.Sprintf(global.REDIS_PHONE, phone), 0, time.Second*60).Err()
	if err != nil {
		zlog.CtxErrorf(ctx, "Restrict multiple access err: %v", err)
	}
	//发送验证码
	if err := PostMessage(text, phone); err != nil {
		zlog.CtxErrorf(ctx, "PostMessage err: %v", err)
	}
	return nil
}
func PostMessage(text string, phone string) error {
	return nil
}
