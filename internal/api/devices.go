package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/global"
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
	if user_id, exists := c.Get(global.TOKEN_USER_ID); exists {
		req.UserId = user_id.(int64)
	}
	if err != nil {
		return
	}
	zlog.CtxInfof(ctx, "ShowDevices request: %v", req)
	resp, err := logic.NewDevicesLogic().ShowDevices(ctx, req)
	if err != nil {
		response.Response(c, nil, err)
	}
	if token, exists := c.Get(global.AUTH_ENUMS_ATOKEN); exists {
		resp.Token = token.(string)
		response.Response(c, resp, err)
	}
	return
}
