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
	// 获取用户飞书open_id
	openID, err := repo.NewFeiShuRepo(global.DB).GetFeiShuOpenID(UserID)
	if err != nil {
		zlog.Errorf("get feishu open_id error:%v", err)
		return
	}
	//fmt.Println(openID)

	// 获取多维表格数据
	resp, err = handler.GetFeiShuList(ctx, openID, ForceUpdate)
	if err != nil {
		zlog.Errorf("get feishu list error:%v", err)
		err = response.ErrResp(err, response.FEISHU_ERROR)
		return
	}

	if err != nil {
		zlog.Errorf("get feishulist error:%v", err)
		return
	} else {
		zlog.Infof("get feishulist success")
	}

	return
}
