package logic

import (
	"context"
	"tgwp/internal/handler"
	"tgwp/internal/types"
	"tgwp/log/zlog"
	"tgwp/util"
	"time"
)

type CodeLogic struct {
}

func NewCodeLogic() *CodeLogic {
	return &CodeLogic{}
}

func (l *CodeLogic) GenCode(ctx context.Context, req types.PhoneReq) (err error) {
	defer util.RecordTime(time.Now())()
	//..... some logic
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
	//这里只是做了简单处理，后期得改进FullToken函数
	resp.ServiceId = 123456
	if req.AutoLogin {
		resp.Atoken, err = util.GenToken(util.FullToken("atoken"))
		resp.Rtoken, err = util.GenToken(util.FullToken("rtoken"))
	} else {
		resp.Atoken, err = util.GenToken(util.FullToken("atoken"))
	}
	return
}
