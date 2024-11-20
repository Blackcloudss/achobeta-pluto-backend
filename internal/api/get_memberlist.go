package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

// GetTeamMemberlist
//
//	@Description: 查询团队列表--用户基础信息
//	@param c
func GetTeamMemberlist(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)

	req, err := types.BindReq[types.MemberlistReq](c)
	if err != nil {
		return
	}
	zlog.CtxInfof(ctx, "GetTeamMemberlist request: %v", req)
	resp, err := logic.NewMemberListic().GetMemberList(ctx, req)
	response.Response(c, resp, err)

	return
}
