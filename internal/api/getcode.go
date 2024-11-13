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

	resp, err := logic.NewCodeLogic().GenLoginData(ctx, req)
	if err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	} else {
		resp.Ip = c.ClientIP()
		resp.UserAgent = c.Request.UserAgent()
		response.NewResponse(c).Success(resp)
	}
	return
}
