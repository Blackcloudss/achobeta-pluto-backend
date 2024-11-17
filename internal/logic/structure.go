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

type StructureLogic struct {
}

func NewStructureLogic() *StructureLogic {
	return &StructureLogic{}
}

// StructureLogic
//
//	@Description:
//	@receiver l
//	@param ctx
//	@param req
//	@return types.TeamStructResp
//	@return error
//
// 获取完整团队架构
func (l *StructureLogic) StructureLogic(ctx context.Context, req types.TeamStructReq) (types.TeamStructResp, error) {
	defer util.RecordTime(time.Now())()

	teamStructures := []types.TeamStructure{}

	// 递归获取节点信息
	err := l.getStructure(ctx, global.ROOT_ID, req.TeamId, &teamStructures)
	if err != nil {
		zlog.CtxErrorf(ctx, "Failed to get children for fatherid: %d, teamid: %d, error: %v", global.ROOT_ID, req.TeamId, err)
		return types.TeamStructResp{}, err
	}

	return types.TeamStructResp{TeamStructures: teamStructures}, nil
}

// getStructure
//
//	@Description:
//	@receiver l
//	@param ctx
//	@param fatherid
//	@param teamid
//	@param result
//	@return error
//
// 递归获取节点的所有子节点
func (l *StructureLogic) getStructure(ctx context.Context, fatherid, teamid int64, result *[]types.TeamStructure) error {
	// 获取当前节点的所有子节点
	children, err := repo.NewStructureRepo(global.DB).GetNode(fatherid, teamid)
	if err != nil {
		return err
	}

	for _, child := range children {
		node := types.TeamStructure{
			TeamId:    teamid,
			MyselfId:  child.MyselfId,
			FatherId:  fatherid,
			NodeName:  child.NodeName,
			IsDeleted: false, // 假设在查询时过滤了已删除的节点
		}

		*result = append(*result, node)

		// 递归获取子节点的子节点
		err = l.getStructure(ctx, child.MyselfId, teamid, result)
		if err != nil {
			zlog.CtxErrorf(ctx, "Failed to get children for fatherid: %d, teamid: %d, error: %v", fatherid, teamid, err)
			return err
		}
	}
	return nil
}
