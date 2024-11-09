package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

// only for test
// 团队信息获取
func GetTeamStructure(c *gin.Context) {

}

func Test(c *gin.Context) {
	// always in the first
	ctx := zlog.GetCtxFromGin(c)

	req, err := types.BindReq[types.TestO1Req](c)
	if err != nil {
		return
	}

	resp, err := logic.NewTestLogic().TestLogic(ctx, req)
	if err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	} else {
		response.NewResponse(c).Success(resp)
	}

	return
}
