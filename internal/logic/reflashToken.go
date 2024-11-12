package logic

import (
	"context"
	"tgwp/global"
	"tgwp/internal/types"
	"tgwp/log/zlog"
	"tgwp/util"
	"time"
)

type TokenLogic struct {
}

func NewTokenLogic() *TokenLogic {
	return &TokenLogic{}
}

func (l *TokenLogic) AtokenLogic(ctx context.Context, data util.TokenData) (resp types.TokenResp, err error) {
	defer util.RecordTime(time.Now())()
	resp.Atoken, err = util.GenToken(data)
	if err != nil {
		zlog.CtxErrorf(ctx, "AtokenLogic err:%v", err)
		return
	}
	return
}

func (l *TokenLogic) RtokenLogic(ctx context.Context, data util.TokenData) (resp types.TokenResp, err error) {
	defer util.RecordTime(time.Now())()
	resp.Rtoken, err = util.GenToken(data)
	if err != nil {
		zlog.CtxErrorf(ctx, "TokenLogic err:%v", err)
	}
	data.Time = global.ATOKEN_EFFECTIVE_TIME
	data.Class = global.AUTH_ENUMS_ATOKEN
	resp.Atoken, err = util.GenToken(data)
	if err != nil {
		zlog.CtxErrorf(ctx, "RtokenLogic err:%v", err)
		return
	}
	return
}
