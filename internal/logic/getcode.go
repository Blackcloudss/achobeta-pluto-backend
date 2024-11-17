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

// GenCode
//
//	@Description: 生成验证码，发送到用户手机
//	@receiver l
//	@param ctx
//	@param req
//	@return err
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

// GenLoginData
//
//	@Description: 为登陆后的用户授予一些个人信息
//	@receiver l
//	@param ctx
//	@param AutoLogin
//	@param resp
//	@return err
func (l *CodeLogic) GenLoginData(ctx context.Context, AutoLogin bool, resp *types.PhoneResp) (err error) {
	defer util.RecordTime(time.Now())()
	node, err := snowflake.NewNode(global.DEFAULT_NODE_ID)
	if err != nil {
		zlog.CtxErrorf(ctx, "NewNode err: %v", err)
		return
	}
	resp.LoginId = snowflake.GenId(node)
	user_id := snowflake.GenId(node)
	if AutoLogin {
		issuer := snowflake.GenId(node)
		resp.Atoken, err = util.GenToken(util.FullToken(global.AUTH_ENUMS_ATOKEN, issuer, user_id))
		resp.Rtoken, err = util.GenToken(util.FullToken(global.AUTH_ENUMS_RTOKEN, issuer, user_id))
		//将点了自动登录的用户的login_id,issuer插入签名表
		data := repo.CommonData{
			UserId:     user_id,
			Issuer:     issuer,
			OnlineTime: time.Now(),
			LoginId:    resp.LoginId,
			IP:         resp.Ip,
			UserAgent:  resp.UserAgent,
		}
		err = repo.NewSignRepo(global.DB).InsertSign(data)
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

// InsertData
//
//	@Description: 获取用户的ip和useragent
//	@param resp
//	@param ip
//	@param user_agent
func InsertData(resp *types.PhoneResp, ip, user_agent string) {
	resp.Ip = ip
	resp.UserAgent = user_agent
}
