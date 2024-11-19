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

type TeamNodeLogic struct {
}

func NewTeamNodeLogic() *TeamNodeLogic {
	return &TeamNodeLogic{}
}

// TeamNodeLogic
//
//	@Description:  保存新增的节点，删除被删除的节点
//	@receiver l
//	@param ctx
//	@param req
//	@return types.PutTeamNodeResp
//	@return error
func (l *TeamNodeLogic) TeamNodeLogic(ctx context.Context, req types.PutTeamNodeReq) (types.PutTeamNodeResp, error) {
	defer util.RecordTime(time.Now())()

	for _, Node := range req.TeamStructures {
		if Node.IsDeleted == false {
			// 没被删除且没有自身节点值 ：新增节点
			err := repo.NewStructureRepo(global.DB).InsertNode(Node)
			if err != nil {
				zlog.CtxErrorf(ctx, "Failed to insert node with NodeName %s: %v", Node.NodeName, err)
				return types.PutTeamNodeResp{}, err
			}
		} else {
			// true ：被删除的节点
			err := repo.NewStructureRepo(global.DB).DeleteNode(Node)
			if err != nil {
				zlog.CtxErrorf(ctx, "Failed to delete node with ID %d: %v", Node.MyselfId, err)
				return types.PutTeamNodeResp{}, err
			}
		}
	}

	return types.PutTeamNodeResp{}, nil
}
