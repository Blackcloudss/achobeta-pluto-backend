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

// 获取完整团队架构
func (l *StructureLogic) StructureLogic(ctx context.Context, req types.TeamStructReq) (types.TeamStructResp, error) {
	defer util.RecordTime(time.Now())()

	var ROOT_ID int64 = 1 // 根节点 ID
	teamStructures := []types.TeamStructure{}

	// 递归获取节点信息
	err := l.getStructure(ctx, ROOT_ID, req.TeamId, &teamStructures)
	if err != nil {
		zlog.CtxErrorf(ctx, "%v", err)
		return types.TeamStructResp{}, err
	}

	return types.TeamStructResp{TeamStructures: teamStructures}, nil
}

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
			zlog.CtxErrorf(ctx, "%v", err)
			return err
		}
	}
	return nil
}

type TeamNodeLogic struct {
}

func NewTeamNodeLogic() *TeamNodeLogic {
	return &TeamNodeLogic{}
}

// 保存新增的节点，删除被删除的节点
func (l *TeamNodeLogic) TeamNodeLogic(ctx context.Context, req types.PutTeamNodeReq) (types.PutTeamNodeResp, error) {
	defer util.RecordTime(time.Now())()

	for _, Node := range req.TeamStructures {
		if Node.IsDeleted == false {
			// 没被删除且没有自身节点值 ：新增节点
			err := repo.NewStructureRepo(global.DB).InsertNode(Node)
			if err != nil {
				zlog.CtxErrorf(ctx, "%v", err)
				return types.PutTeamNodeResp{}, err
			}
		} else {
			// true ：被删除的节点
			err := repo.NewStructureRepo(global.DB).DeleteNode(Node)
			if err != nil {
				zlog.CtxErrorf(ctx, "%v", err)
				return types.PutTeamNodeResp{}, err
			}
		}
	}

	return types.PutTeamNodeResp{}, nil
}
