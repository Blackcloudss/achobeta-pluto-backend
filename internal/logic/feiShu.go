package logic

import (
	"context"
	"tgwp/global"
	"tgwp/internal/handler"
	"tgwp/internal/repo"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

type FeiShuLogic struct {
}

func NewFeiShuLogic() *FeiShuLogic {
	return &FeiShuLogic{}
}

// GetFeiShuList 获取飞书多维表格
func (l *FeiShuLogic) GetFeiShuList(ctx context.Context, UserID int64, ForceUpdate bool) (resp types.GetFeiShuListResp, err error) {
	// 1.获取用户飞书open_id
	openID, err := repo.NewFeiShuRepo(global.DB).GetFeiShuOpenID(UserID)
	if err != nil {
		zlog.Errorf("get feishu open_id error:%v", err)
		err = response.ErrResp(err, response.FEISHU_ERROR)
		return
	}
	// 2.先检查是否需要更新
	needUpdate, err := handler.CheckUpdate(ctx, ForceUpdate)
	if err != nil {
		zlog.Errorf("check update error:%v", err)
		err = response.ErrResp(err, response.FEISHU_ERROR)
		return
	}
	// 3.如果需要更新，则更新
	if needUpdate {
		err = handler.UpdateFeiShuList(ctx)
	}
	// 4.获取列表数据
	resp, err = handler.GetFeiShuList(ctx, openID)
	if err != nil {
		zlog.Errorf("get feishu list error:%v", err)
		err = response.ErrResp(err, response.FEISHU_ERROR)
		return
	}
	return
}
