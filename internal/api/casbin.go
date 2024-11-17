package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/global"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

// GetPower
//
//	@Description:
//	@param c
func GetPower(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)
	//这里重新 set是因为  global.TOKEN_USER_ID = "UserId"
	//和 表的 types.RuleReq 的 UserId int64 `json:"user_id" ` 不一致，所有重新 Set
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
