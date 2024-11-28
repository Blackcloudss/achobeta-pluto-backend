package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

// ShowDevices
//
//	@Description: 展示常用设备
//	@param c
func ShowDevices(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)
	req, err := types.BindReq[types.DevicesReq](c)
	if err != nil {
		return
	}
	req.UserId = logic.GetUserId(c)
	zlog.CtxInfof(ctx, "ShowDevices request: %v", req)
	resp, err := logic.NewDevicesLogic().ShowDevices(ctx, req)
	response.Response(c, resp, err)
	return
}

// ModifyDeviceName
//
//	@Description: 修改设备名称
//	@param c
func ModifyDeviceName(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)
	req, err := types.BindReq[types.ModifyDeviceNameReq](c)
	if err != nil {
		return
	}
	zlog.CtxInfof(ctx, "ShowDevices request: %v", req)
	err = logic.NewDevicesLogic().ModifyDeviceName(ctx, req)
	response.Response(c, nil, err)
	return
}
