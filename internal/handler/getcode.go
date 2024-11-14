package handler

import (
	"context"
	"errors"
	"fmt"
	"tgwp/global"
	"tgwp/log/zlog"
	"tgwp/util"
	"time"
)

// 用作验证码处理
func PostCode(ctx context.Context, phone string) error {
	if !AccessCode(ctx, phone) {
		zlog.CtxErrorf(ctx, "Access code error")
		return errors.New("该手机号在一分钟内已经发送给验证码")
	}
	//生成6位数字的验证码
	code := util.RandomCode()
	text := fmt.Sprintf("你的验证码是%s", code)
	//将验证码放入redis5分钟
	err := global.Rdb.Set(ctx, fmt.Sprintf(global.REDIS_PHONE_CODE, phone), code, time.Second*300).Err()
	if err != nil {
		zlog.CtxErrorf(ctx, "Store the verification code err: %v", err)
	}
	//防刷处理
	err = global.Rdb.Set(ctx, fmt.Sprintf(global.REDIS_PHONE, phone), 0, time.Second*60).Err()
	if err != nil {
		zlog.CtxErrorf(ctx, "Restrict multiple access err: %v", err)
		return err
	}
	//发送验证码到用户手机
	if err := PostMessage(text, phone); err != nil {
		zlog.CtxErrorf(ctx, "PostMessage err: %v", err)
	}
	return nil
}
func PostMessage(text string, phone string) error {
	return nil
}
func CompareCode(ctx context.Context, code, phone string) bool {
	return code == global.Rdb.Get(ctx, fmt.Sprintf(global.REDIS_PHONE_CODE, phone)).Val()
}
func AccessCode(ctx context.Context, phone string) bool {
	//存在，即手机号一分钟内被记录
	if exist := global.Rdb.Exists(ctx, fmt.Sprintf(global.REDIS_PHONE, phone)).Val(); exist == 1 {
		return false
	}
	return true
}
