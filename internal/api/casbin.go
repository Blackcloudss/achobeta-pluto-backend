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

	if req.UserId == "Null" {
		response.NewResponse(c).Error(response.PARAM_IS_BLANK)
		return
	}

	if req.TeamId == "NULL" {
		// null 时 仅返回第一个团队ID，所有的团队ID，状态码code，信息获取（成功/失败）msg
	}

	urls, err := repo.Getcasbin(req)
	if err != nil {

	}
	//把找出来的url给出参data
	resq.Data = urls

	//获取团队id
	FTeamID, TeamID, errs := repo.GetTeamId()
	if errs != nil {

	}
	resq.FirstTeamID = FTeamID
	resq.TeamID = TeamID
}
