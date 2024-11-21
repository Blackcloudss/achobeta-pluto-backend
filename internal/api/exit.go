package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/global"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

// ExitSystem
//
//	@Description: 用户自己退出接口
//	@param c
func ExitSystem(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)
	req, err := types.BindJson[types.TokenReq](c)
	if err != nil {
		return
	}
	zlog.CtxInfof(ctx, "ExitSystem request: %v", req)
	err = logic.NewExitLogic().ExitSystem(ctx, req)
	response.Response(c, nil, err)
}

func RemoveDevice(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)
	req, err := types.BindJson[types.RemoveDeviceReq](c)
	if err != nil {
		return
	}
	zlog.CtxInfof(ctx, "RemoveDevice request: %v", req)
	err = logic.NewDevicesLogic().RemoveDevices(ctx, req)
	if err != nil {
		response.Response(c, nil, err)
	}
	if token, exists := c.Get(global.AUTH_ENUMS_ATOKEN); exists {
		response.NewResponse(c).Success(token)
	}
}
