package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/response"
)

func GetTeamStructure(c *gin.Context) {

	ts, exists := c.Get("team_id")
	if !exists {
		//‘teamid’的值不存在
		response.NewResponse(c).Error(response.PARAM_IS_BLANK)
		return
	}
}
