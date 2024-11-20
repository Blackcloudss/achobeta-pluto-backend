package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"tgwp/global"
	"tgwp/internal/repo"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
	"tgwp/util"
	"time"
)

type CreateMemberLogic struct{}

func NewCreateMemberLogic() *CreateMemberLogic {
	return &CreateMemberLogic{}
}

// CreateMember
//
//	@Description: 新增团队成员
//	@receiver l
//	@param ctx
//	@param req
//	@return *types.CreateMembersResp
//	@return error
func (l *CreateMemberLogic) CreateMember(ctx context.Context, req types.CreateMemberReq) (*types.CreateMembersResp, error) {
	defer util.RecordTime(time.Now())()

	err := repo.NewCreateMemberRepo(global.DB).CreateMember(req)
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidData) {
			zlog.CtxWarnf(ctx, "Invalid Data: %v", err)
			return nil, response.ErrResp(err, InvalidData)
		} else {
			zlog.Errorf("get user error: %v", err)
			return nil, response.ErrResp(err, codeUserFoundField)
		}

	}
	return &types.CreateMembersResp{}, nil

}

type DeleteMemberLogic struct{}

func NewDeleteMemberLogic() *DeleteMemberLogic {
	return &DeleteMemberLogic{}
}

// DeleteMember
//
//	@Description: 删除该成员在这个团队的关系
//	@receiver l
//	@param ctx
//	@param req
//	@return *types.DeleteMembersResp
//	@return error
func (l *DeleteMemberLogic) DeleteMember(ctx context.Context, req types.DeleteMemberReq) (*types.DeleteMembersResp, error) {
	defer util.RecordTime(time.Now())()

	err := repo.NewDeleteMemberRepo(global.DB).DeleteMember(req.MemberId, req.TeamId)
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidData) {
			zlog.CtxWarnf(ctx, "Invalid Data: %v", err)
			return nil, response.ErrResp(err, InvalidData)
		} else {
			zlog.Errorf("get user error: %v", err)
			return nil, response.ErrResp(err, codeUserFoundField)
		}

	}
	return &types.DeleteMembersResp{}, nil

}

// 编辑成员信息
type PutMemberLogic struct{}

func NewPutMemberLogic() *PutMemberLogic {
	return &PutMemberLogic{}
}

// PutMember
//
//	@Description: 更改成员信息
//	@receiver l
//	@param ctx
//	@param req
//	@return *types.PutTeamMemberResp
//	@return error
func (l *PutMemberLogic) PutMember(ctx context.Context, req types.PutTeamMemberReq) (*types.PutTeamMemberResp, error) {
	defer util.RecordTime(time.Now())()

	err := repo.NewPutMemberRepo(global.DB).PutMember(req)
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidData) {
			zlog.CtxWarnf(ctx, "Invalid Data: %v", err)
			return nil, response.ErrResp(err, InvalidData)
		} else {
			zlog.Errorf("get user error: %v", err)
			return nil, response.ErrResp(err, codeUserFoundField)
		}

	}
	return &types.PutTeamMemberResp{}, nil
}
