package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

// CreateTeam
//
//	@Description: 新增团队
//	@param c
func CreateTeam(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)
	req, err := types.BindReq[types.CreateTeamReq](c)
	if err != nil {
		zlog.CtxErrorf(ctx, "PostTeam err:%v", err)
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	}
	zlog.CtxInfof(ctx, "PostTeam request: %v", req)
	resp, err := logic.NewTeamLogic().TeamLogic(ctx, req)

	response.Response(c, resp, err)

	return

}
