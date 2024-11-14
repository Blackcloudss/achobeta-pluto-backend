package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/global"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

func GetPower(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)
	if userid, exists := c.Get(global.TOKEN_USER_ID); exists {
		// 让userid直接和结构体的user_id绑定
		c.Set("user_id", userid)
	}

	req, err := types.BindReq[types.RuleReq](c)

	if err != nil {
		zlog.CtxErrorf(ctx, "GetPower err:%v", err)
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	}
	zlog.CtxInfof(ctx, "GetPower request: %v", req)
	//获取出参
	resp, err := logic.NewCasbinLogic().CasbinLogic(ctx, req)

	if err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	} else {
		response.NewResponse(c).Success(resp)
	}
	return
}
