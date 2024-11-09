package logic

import (
	"context"
	"tgwp/internal/repo"
	"tgwp/internal/types"
	"tgwp/log/zlog"
	"tgwp/util"
	"time"
)

type TestLogic struct {
}

func NewTestLogic() *TestLogic {
	return &TestLogic{}
}
func (l *TestLogic) TestLogic(ctx context.Context, req types.TestO1Req) (resp types.Test01Resp, err error) {
	defer util.RecordTime(time.Now())()
	//..... some logic

	user, err := repo.NewTestRepo().GetUserById(req.UserID)
	if err != nil {
		zlog.CtxErrorf(ctx, "%v", err)
	}
	resp.Name = user.Name
	resp.Age = user.Age
	return
}
