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

// CasbinLogic  获取权限组
func (l *CasbinLogic) CasbinLogic(ctx context.Context, req types.RuleReq) (resp types.RuleResp, err error) {
	defer util.RecordTime(time.Now())()
	// 前端没有传团队id 时 仅返回第一个团队ID，所有的团队ID，状态码code，信息获取（成功/失败）msg
	if req.TeamId != 0 {
		//获取 该用户在该团队所拥有的 特殊 url
		urls, err := repo.NewCasbinRepo(global.DB).GetCasbin(req.UserId, req.TeamId)
		if err != nil {
			zlog.CtxErrorf(ctx, "%v", err)
			return
		}
		//把找出来的url给出参 Url
		resp.Url = urls

		//获取 该用户在该团队所拥有的 权限级别
		Level, err := repo.NewLevelRepo(global.DB).GetLevel(req.UserId, req.TeamId)
		if err != nil {
			zlog.CtxErrorf(ctx, "%v", err)
			return
		}
		resp.Level = Level
	}

	//获取 第一个团队ID，所有的团队ID
	FTeamID, TeamID, err := repo.NewTeamIdRepo(global.DB).GetTeamId(req.UserId)
	if err != nil {
		zlog.CtxErrorf(ctx, "%v", err)
		return
	}
	resp.FirstTeamID = FTeamID
	resp.TeamID = TeamID
	return
}
