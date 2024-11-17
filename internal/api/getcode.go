package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/handler"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
	"tgwp/util"
)

// GetCode
//
//	@Description: 获取验证码
//	@param c
func GetCode(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)
	req, err := types.BindReq[types.PhoneReq](c)
	if err != nil {
		return
	}
	//校验手机号
	if flag := util.IndetifyPhone(req.Phone); !flag {
		response.NewResponse(c).Error(response.PHONE_ERROR)
		return
	}
	zlog.CtxInfof(ctx, "GetCode request: %v", req)
	err = logic.NewCodeLogic().GenCode(ctx, req)
	if err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	} else {
		response.NewResponse(c).Success(response.SUCCESS)
	}
	return
}

// LoginWithCode
//
//	@Description: 用验证码登录
//	@param c
func LoginWithCode(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)
	req, err := types.BindReq[types.PhoneReq](c)
	if err != nil {
		return
	}
	zlog.CtxInfof(ctx, "LoginWithCode request: %v", req)
	if !handler.CompareCode(ctx, req.Code, req.Phone) {
		response.NewResponse(c).Error(response.CAPTCHA_ERROR)
		return
	}
	var resp types.PhoneResp
	logic.InsertData(&resp, c.ClientIP(), c.Request.UserAgent())
	err = logic.NewCodeLogic().GenLoginData(ctx, req.AutoLogin, &resp)
	if err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	} else {
		response.NewResponse(c).Success(resp)
	}
	return
}
