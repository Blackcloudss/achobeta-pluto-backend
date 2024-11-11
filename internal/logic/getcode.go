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
	//这里只是做了简单处理，后期得改进FullToken函数
	resp.Atoken, err = util.GenToken(util.FullToken("atoken"))
	resp.Rtoken, err = util.GenToken(util.FullToken("rtoken"))
	return
}
