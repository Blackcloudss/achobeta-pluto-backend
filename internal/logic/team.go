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

//	TeamLogic
//	@Description:
//
// 新增团队 ，初始化团队架构
type TeamLogic struct {
}

func NewTeamLogic() *TeamLogic {
	return &TeamLogic{}
}

func (l *TeamLogic) TeamLogic(ctx context.Context, req types.PostTeamReq) (resp types.PostTeamResp, err error) {
	defer util.RecordTime(time.Now())()
	resp, err = repo.NewPostTeamRepo(global.DB).PostTeam(req.Name)
	if err != nil {
		zlog.CtxErrorf(ctx, "v", err)
	}
	return
}
