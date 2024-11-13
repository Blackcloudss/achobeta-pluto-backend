package logic

import (
	"context"
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

func (l *StructureLogic) StructureLogic(ctx context.Context, req types.TeamStructReq) (resp types.TeamStructResp, err error) {
	defer util.RecordTime(time.Now())()

	//user, err := repo.NewTestRepo().GetUserById(req.UserID)
	//待开发
	if err != nil {
		zlog.CtxErrorf(ctx, "%v", err)
	}

	return
}
