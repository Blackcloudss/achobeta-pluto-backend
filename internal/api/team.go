package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

//	 PostTeam
//
//		@Description:
//		@param c
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

	if err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	} else {
		response.NewResponse(c).Success(resp)
	}

	return

}
