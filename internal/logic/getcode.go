package logic

import (
	"context"
	"tgwp/internal/types"
	"tgwp/util"
	"time"
)

type CodeLogic struct {
}

func NewCodeLogic() *CodeLogic {
	return &CodeLogic{}
}

// TestLogic 逻辑层 用做逻辑处理相关操作
func (l *CodeLogic) CodeLogic(ctx context.Context, req types.PhoneReq) (resp types.PhoneResp, err error) {
	defer util.RecordTime(time.Now())()
	//..... some logic
	//暂时不处理redis层面，直接让验证码为123456
	resp.Atoken = "暂无"
	resp.Rtoken = "暂无"
	return
}
