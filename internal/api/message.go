package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

// SetMessage 设置消息, 存入 [消息表]
func SetMessage(c *gin.Context) {
	// 解析请求参数
	ctx := zlog.GetCtxFromGin(c)
	req, err := types.BindReq[types.SetMessageReq](c)
	if err != nil {
		return
	}
	zlog.CtxInfof(ctx, "Casbin request: %v", req)

	// logic 层处理
	resp, err := logic.NewMessageLogic().SetMessage(c, req)

	// 响应
	if err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	} else {
		response.NewResponse(c).Success(resp)
	}
	return
}

// JoinMessage 连接用户与消息, 存入 [用户-消息表]
func JoinMessage(c *gin.Context) {
	// 解析请求参数
	ctx := zlog.GetCtxFromGin(c)
	req, err := types.BindReq[types.JoinMessageReq](c)
	if err != nil {
		return
	}
	zlog.CtxInfof(ctx, "Casbin request: %v", req)
	//fmt.Println(util.GenToken(util.TokenData{Userid: "114514", Class: "1", Issuer: "1", Time: time.Hour * 24 * 365}))

	// logic 层处理
	resp, err := logic.NewMessageLogic().JoinMessage(c, req)

	// 响应
	if err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	} else {
		response.NewResponse(c).Success(resp)
	}

	return
}
