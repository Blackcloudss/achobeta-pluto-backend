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

type CasbinLogic struct {
}

func NewCasbinLogic() *CasbinLogic {
	return &CasbinLogic{}
}

func (l *CasbinLogic) CasbinLogic(ctx context.Context, req types.RuleReq) (resp types.RuleResp, err error) {
	defer util.RecordTime(time.Now())()

	if req.TeamId == 0 {
		// 前端没有传团队id 时 仅返回第一个团队ID，所有的团队ID，状态码code，信息获取（成功/失败）msg
	}

	urls, err := repo.NewCasbinRepo(global.DB).Getcasbin(req.UserId, req.TeamId)
	if err != nil {
		zlog.CtxErrorf(ctx, "%v", err)

	}
	//把找出来的url给出参data
	resp.Url = urls

	//获取团队id
	//不用传参
	FTeamID, TeamID, errs := repo.NewTeamIdRepo(global.DB).GetTeamId()
	if errs != nil {
		zlog.CtxErrorf(ctx, "%v", err)
		return
	}
	resp.FirstTeamID = FTeamID
	resp.TeamID = TeamID
	//正在开发

	return
}
