package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

// CreateTeamMember
//
//	@Description: 新增用户
//	@param c
func CreateTeamMember(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)

	req, err := types.BindReq[types.CreateMemberReq](c)
	if err != nil {
		return
	}
	zlog.CtxInfof(ctx, "Create TeamMember request: %v", req)
	resp, err := logic.NewCreateMemberLogic().CreateMember(ctx, req)
	// 更加人性化的response返回，这样减少重复代码的书写
	response.Response(c, resp, err)

	return
}

// DeleteTeamMember
//
//	@Description: 删除用户
//	@param c
func DeleteTeamMember(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)

	req, err := types.BindReq[types.DeleteMemberReq](c)
	if err != nil {
		return
	}
	zlog.CtxInfof(ctx, "DeleteTeamMember request: %v", req)
	resp, err := logic.NewDeleteMemberLogic().DeleteMember(ctx, req)
	response.Response(c, resp, err)

	return
}
