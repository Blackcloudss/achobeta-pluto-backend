package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

func GetTeamStructure(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)
	req, err := types.BindReq[types.TeamStructReq](c)
	if err != nil {
		zlog.CtxErrorf(ctx, "GetTeamStructure err:%v", err)
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	}
	zlog.CtxInfof(ctx, "GetTeamStructure request: %v", req)
	resp, err := logic.NewStructureLogic().StructureLogic(ctx, req)

	if err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	} else {
		response.NewResponse(c).Success(resp)
	}

	return
}

func PutTeamNode(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)
	req, err := types.BindReq[types.PutTeamNodeReq](c)
	if err != nil {
		zlog.CtxErrorf(ctx, "PutTeamNode err:%v", err)
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	}
	zlog.CtxInfof(ctx, "PutTeamNode request: %v", req)
	resp, err := logic.NewTeamNodeLogic().TeamNodeLogic(ctx, req)

	if err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	} else {
		response.NewResponse(c).Success(resp)
	}

	return
}
