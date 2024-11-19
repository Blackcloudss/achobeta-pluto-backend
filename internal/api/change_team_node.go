package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

// PutTeamNode
//
//	@Description: 保存 更改了的节点信息
//	@param c
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

	response.Response(c, resp, err)
	return
}
