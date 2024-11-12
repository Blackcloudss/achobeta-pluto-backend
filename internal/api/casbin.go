package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/repo"
	"tgwp/internal/response"
	"tgwp/internal/types"
)

func GetPower(c *gin.Context) {
	var req types.RuleReq
	var resq types.RuleResq

	c.ShouldBind(&req)

	if req.UserId == 0 {
		// 前端没有传用户id 时
		response.NewResponse(c).Error(response.PARAM_NOT_COMPLETE)
		return
	}

	if req.TeamId == 0 {
		// 前端没有传团队id 时 仅返回第一个团队ID，所有的团队ID，状态码code，信息获取（成功/失败）msg
	}

	urls, err := repo.NewCasbinRepo().Getcasbin(req.UserId, req.TeamId)
	if err != nil {
		response.NewResponse(c).Error(response.PARAM_IS_BLANK)
		return
	}
	//把找出来的url给出参data
	resq.Data = urls

	//获取团队id
	FTeamID, TeamID, errs := repo.GetTeamId()
	if errs != nil {
		response.NewResponse(c).Error(response.PARAM_IS_BLANK)
		return
	}
	resq.FirstTeamID = FTeamID
	resq.TeamID = TeamID

	//还在开发
	return
}
