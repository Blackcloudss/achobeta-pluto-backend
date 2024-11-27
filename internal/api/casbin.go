package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/handler"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

// GetPower
//
//	@Description: 获得权限组
//	@param c
func GetPower(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)

	//正式使用，测试时需注释掉
	UserId := handler.GetUserId(c)
	req, err := types.BindReq[types.RuleReq](c)

	if err != nil {
		zlog.CtxErrorf(ctx, "GetPower err:%v", err)
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	}
	zlog.CtxInfof(ctx, "GetPower request: %v", req)
	//获取出参
	resp, err := logic.NewCasbinLogic().GetCasbin(ctx, UserId, req.TeamId) //正式使用
	//resp, err := logic.NewCasbinLogic().GetCasbin(ctx, req.UserId, req.TeamId) //测试时使用
	response.Response(c, resp, err)
	return
}
