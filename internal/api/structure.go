package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/types"
)

func GetTeamStructure(c *gin.Context) {
	var req types.Team_StructReq
	var resq types.TeamStructResq

	c.ShouldBind(&req)

	//开发中

	return

}
