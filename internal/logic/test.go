package logic

import (
	"context"
	"tgwp/internal/repo"
	"tgwp/internal/response"
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

// 这里定义我们内部logic的错误（非公有的常见类型的错误）
var (
	codeUserFoundField = response.MsgCode{Code: 40013, Msg: "用户查询失败"}
)

// TestLogic 逻辑层 用做逻辑处理相关操作
func (l *TestLogic) TestLogic(ctx context.Context, req types.TestO1Req) (resp *types.Test01Resp, err error) {
	defer util.RecordTime(time.Now())()
	//..... some logic

	user, err := repo.NewTestRepo().GetUserById(req.UserID)
	if err != nil {
		zlog.Errorf("get user error: %v", err)
		//注意 codeUserFoundField 只是事例，具体根据实际情况定义 response.ErrResp这里用作包装错误和响应，使得错误更加通用
		return nil, response.ErrResp(err, codeUserFoundField)
	}
	resp.Name = user.Name
	resp.Age = user.Age

	return
}
