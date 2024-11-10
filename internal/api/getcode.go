package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
	"tgwp/util"
)

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
	_, err = logic.NewCodeLogic().CodeLogic(ctx, req)
	if err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	} else {
		response.NewResponse(c).Success(response.SUCCESS)
	}
	return
}
func LoginWithCode(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)
	req, err := types.BindReq[types.PhoneReq](c)
	if err != nil {
		return
	}
	zlog.CtxInfof(ctx, "LoginWithCode request: %v", req)
	//暂时处理，若验证码不为123456，则响应错误
	if req.Code != "123456" {
		response.NewResponse(c).Error(response.CAPTCHA_ERROR)
		return
	}

	resp, err := logic.NewCodeLogic().CodeLogic(ctx, req)
	if err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	} else {
		response.NewResponse(c).Success(resp)
	}
	return
}
