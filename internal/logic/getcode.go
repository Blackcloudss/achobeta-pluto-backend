package logic

import (
	"context"
	"tgwp/internal/handler"
	"tgwp/internal/types"
	"tgwp/log/zlog"
	"tgwp/util"
	"tgwp/util/snowflake"
	"time"
)

type CodeLogic struct {
}

func NewCodeLogic() *CodeLogic {
	return &CodeLogic{}
}

func (l *CodeLogic) GenCode(ctx context.Context, req types.PhoneReq) (err error) {
	defer util.RecordTime(time.Now())()
	//生成随机验证码并发送到对应用户
	err = handler.PostCode(ctx, req.Phone)
	if err != nil {
		zlog.CtxErrorf(ctx, "GenCode err: %v", err)
		return
	}
	return
}
func (l *CodeLogic) GenLoginData(ctx context.Context, req types.PhoneReq) (resp types.PhoneResp, err error) {
	defer util.RecordTime(time.Now())()
	//传入不同节点是为了生成不同的id,不设置为1是为了区分全局变量
	resp.ServiceId, err = snowflake.GenId(2)
	if err != nil {
		zlog.CtxErrorf(ctx, "GenLogin err: %v", err)
		return
	}
	user_id, err := snowflake.GenId(3)
	if err != nil {
		zlog.CtxErrorf(ctx, "GenLogin err: %v", err)
		return
	}
	issuer, err := snowflake.GenId(4)
	if err != nil {
		zlog.CtxErrorf(ctx, "GenLogin err: %v", err)
		return
	}
	if req.AutoLogin {
		resp.Atoken, err = util.GenToken(util.FullToken("atoken", issuer, user_id))
		resp.Rtoken, err = util.GenToken(util.FullToken("rtoken", issuer, user_id))
	} else {
		resp.Atoken, err = util.GenToken(util.FullToken("atoken", issuer, user_id))
	}
	return
}
