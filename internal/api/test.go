package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

// Test  api层 仅作为校验参数和返回相应，复杂逻辑交给logic层处理
func Test(c *gin.Context) {
	// always in the first
	ctx := zlog.GetCtxFromGin(c)

	req, err := types.BindReq[types.TestO1Req](c)
	if err != nil {
		return
	}
	zlog.CtxInfof(ctx, "Test request: %v", req)
	resp, err := logic.NewTestLogic().TestLogic(ctx, req)

	if err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	} else {
		response.NewResponse(c).Success(resp)
	}

	return
}
