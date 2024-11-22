package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tgwp/internal/handler"
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
	fmt.Println(req.LineNumber, req.PageNumber)
	req.UserId = handler.GetUserId(c)
	zlog.CtxInfof(ctx, "ShowDevices request: %v", req)
	resp, err := logic.NewDevicesLogic().ShowDevices(ctx, req)
	if err != nil {
		response.Response(c, nil, err)
	}
	response.Response(c, resp, err)
	return
}
