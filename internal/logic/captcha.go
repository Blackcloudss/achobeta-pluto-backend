package logic

import (
	"context"
	"tgwp/global"
	"tgwp/internal/handler"
	"tgwp/internal/model"
	"tgwp/internal/repo"
	"tgwp/internal/response"
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
	//校验手机号
	if flag := util.IndetifyPhone(req.Phone); !flag {
		return response.ErrResp(err, response.PHONE_ERROR)
	}
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
func (l *CodeLogic) GenLoginData(ctx context.Context, req types.PhoneReq, ip, user_agent string) (resp types.PhoneResp, err error) {
	defer util.RecordTime(time.Now())()
	node, err := snowflake.NewNode(global.DEFAULT_NODE_ID)
	if err != nil {
		zlog.CtxErrorf(ctx, "NewNode err: %v", err)
		return resp, response.ErrResp(err, response.COMMON_FAIL)
	}
	if !handler.CompareCode(ctx, req.Code, req.Phone) {
		return resp, response.ErrResp(err, response.CAPTCHA_ERROR)
	}
	var user_id int64
	user_id, resp.IsTeam, err = repo.NewMemberRepo(global.DB).JudgeUser(req.Phone)
	if err != nil {
		zlog.CtxErrorf(ctx, "GenLoginData err: %v", err)
		//这里的err只可能是数据库出现错误了
		return resp, response.ErrResp(err, response.COMMON_FAIL)
	}
	if !resp.IsTeam {
		//非团队人员直接返回
		return resp, err
	}
	resp.Ip = ip
	resp.UserAgent = user_agent
	if req.AutoLogin {
		issuer := snowflake.GenId(node)
		resp.Atoken, err = util.GenToken(util.FullToken(global.AUTH_ENUMS_ATOKEN, issuer, user_id))
		resp.Rtoken, err = util.GenToken(util.FullToken(global.AUTH_ENUMS_RTOKEN, issuer, user_id))
		//将点了自动登录的用户的信息插入签名表
		data := model.Sign{
			UserId:     user_id,
			Issuer:     issuer,
			OnlineTime: time.Now(),
			IP:         resp.Ip,
			UserAgent:  resp.UserAgent,
		}
		//由于这个id只是用于移除常用设备，和填充常用设备的名字，所以也是只有自动登陆的有
		resp.Id, err = repo.NewSignRepo(global.DB).InsertSign(data)
		if err != nil {
			zlog.CtxErrorf(ctx, "InsertSign err: %v", err)
			return resp, response.ErrResp(err, response.COMMON_FAIL)
		}
	} else {
		issuer := "" //没有自动登录，做置空处理
		resp.Atoken, err = util.GenToken(util.FullToken(global.AUTH_ENUMS_ATOKEN, issuer, user_id))
	}
	return
}
