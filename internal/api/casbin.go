package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

func GetPower(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)

	req, err := types.BindReq[types.RuleReq](c)
	if err != nil {
		return
	}
	zlog.CtxInfof(ctx, "Casbin request: %v", req)
	resp, err := logic.NewCasbinLogic().CasbinLogic(ctx, req)

	if err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	} else {
		response.NewResponse(c).Success(resp)
	}

	//还在开发
	return
}
