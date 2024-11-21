package logic

import (
	"context"
	"tgwp/global"
	"tgwp/internal/repo"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
	"tgwp/util"
	"time"
)

type DevicesLogic struct {
}

func NewDevicesLogic() *DevicesLogic {
	return &DevicesLogic{}
}

// ShowDevices
//
//	@Description: 展示常用设备
//	@receiver l
//	@param ctx
//	@param data
//	@return resp
//	@return err
func (l *DevicesLogic) ShowDevices(ctx context.Context, req types.DevicesReq) (resp types.DevicesResp, err error) {
	defer util.RecordTime(time.Now())()
	resp, err = repo.NewSignRepo(global.DB).ShowDevices(req)
	if err != nil {
		zlog.CtxErrorf(ctx, "ShowDevices err: %v", err)
		return resp, response.ErrResp(err, response.COMMON_FAIL)
	}
	return
}

// RemoveDevices
//
//	@Description: 移除常用设备
//	@receiver l
//	@param ctx
//	@param req
//	@return resp
//	@return err
func (l *DevicesLogic) RemoveDevices(ctx context.Context, req types.RemoveDeviceReq) (err error) {
	defer util.RecordTime(time.Now())()
	err = repo.NewSignRepo(global.DB).DeleteSignByLoginId(req.LoginId)
	if err != nil {
		zlog.CtxErrorf(ctx, "error deleting sign by loginId: %v", err)
		return response.ErrResp(err, response.COMMON_FAIL)
	}
	return
}
