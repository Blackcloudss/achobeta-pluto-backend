package logic

import (
	"context"
	"tgwp/global"
	"tgwp/internal/handler"
	"tgwp/internal/repo"
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
	node, err := snowflake.NewNode(global.DEFAULT_NODE_ID)
	if err != nil {
		zlog.CtxErrorf(ctx, "NewNode err: %v", err)
		return
	}
	resp.LoginId = snowflake.GenId(node)
	user_id := snowflake.GenId(node)
	if req.AutoLogin {
		issuer := snowflake.GenId(node)
		resp.Atoken, err = util.GenToken(util.FullToken(global.AUTH_ENUMS_ATOKEN, issuer, user_id))
		resp.Rtoken, err = util.GenToken(util.FullToken(global.AUTH_ENUMS_RTOKEN, issuer, user_id))
		//将点了自动登录的用户的login_id,issuer插入签名表
		err = repo.NewSignRepo(global.DB).InsertSign(resp.LoginId, issuer)
		if err != nil {
			zlog.CtxErrorf(ctx, "InsertSign err: %v", err)
			return
		}
	} else {
		issuer := "" //没有自动登录，做置空处理
		resp.Atoken, err = util.GenToken(util.FullToken(global.AUTH_ENUMS_ATOKEN, issuer, user_id))
	}
	return
}
