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

var InvalidData = response.MsgCode{10001, "无效数据"}

type MemberListLogic struct {
}

func NewMemberListic() *MemberListLogic {
	return &MemberListLogic{}
}

// GetMemberList
//
//	@Description: 查看用户基本信息
//	@receiver l
//	@param ctx
//	@param req
//	@return *types.MemberlistResp
//	@return error
func (l *MemberListLogic) GetMemberList(ctx context.Context, req types.MemberlistReq) (*types.MemberlistResp, error) {
	defer util.RecordTime(time.Now())()

	users, err := repo.NewMemberlistRepo(global.DB).MemberlistRepo(req.TeamID, req.Page, req.Perpage)
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidData) {
			zlog.CtxWarnf(ctx, "InvalidData: %v", err)
			return nil, response.ErrResp(err, InvalidData)
		} else {
			zlog.Errorf("get user error: %v", err)
			return nil, response.ErrResp(err, codeUserFoundField)
		}

	}

	return &users, nil
}
