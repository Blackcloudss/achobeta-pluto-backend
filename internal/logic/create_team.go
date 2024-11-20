package logic

import (
	"context"
	"tgwp/global"
	"tgwp/internal/repo"
	"tgwp/internal/types"
	"tgwp/log/zlog"
	"tgwp/util"
	"time"
)

// TeamLogic
// @Description: 新增团队 ，初始化团队架构
type TeamLogic struct {
}

func NewTeamLogic() *TeamLogic {
	return &TeamLogic{}
}

func (l *TeamLogic) TeamLogic(ctx context.Context, req types.CreateTeamReq) (resp types.CreateTeamResp, err error) {
	defer util.RecordTime(time.Now())()
	resp, err = repo.NewCreateTeamRepo(global.DB).CreateTeam(req.Name)
	if err != nil {
		zlog.CtxErrorf(ctx, "新增团队失败：%v", err)
	}
	return
}
