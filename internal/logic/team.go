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

var (
	codeTeamCreateField = response.MsgCode{Code: 40030, Msg: "新增团队失败"}
)

type TeamLogic struct {
}

func NewTeamLogic() *TeamLogic {
	return &TeamLogic{}
}

// CreateTeam
//
//	@Description: 创建团队
//	@receiver l
//	@param ctx
//	@param req
//	@return resp
//	@return err
func (l *TeamLogic) CreateTeam(ctx context.Context, req types.CreateTeamReq) (resp *types.CreateTeamResp, err error) {
	defer util.RecordTime(time.Now())()
	resp, err = repo.NewTeamRepo(global.DB).CreateTeam(req.Name)
	if err != nil {
		zlog.CtxErrorf(ctx, "创建团队失败: %v", err)
		return nil, response.ErrResp(err, codeTeamCreateField)
	}
	return resp, nil
}
