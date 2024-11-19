package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/global"
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
	UserID, _ := c.Get(global.TOKEN_USER_ID)
	zlog.CtxInfof(ctx, "Casbin request: %v", req)

	// logic 层处理
	resp, err := logic.NewMessageLogic().JoinMessage(c, req, UserID.(string))

	// 响应
	if err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	} else {
		response.NewResponse(c).Success(resp)
	}

	return
}

// GetMessage 获取消息, 从 [消息表] 获取
func GetMessage(c *gin.Context) {
	//fmt.Println(util.GenToken(util.TokenData{Userid: "114514", Class: "atoken", Issuer: "1", Time: time.Hour * 24 * 365}))

	// 解析请求参数
	ctx := zlog.GetCtxFromGin(c)
	UserID, _ := c.Get(global.TOKEN_USER_ID)

	pageStr := c.DefaultQuery("page", "1")
	timestampStr := c.DefaultQuery("timestamp", "0")

	zlog.CtxInfof(ctx, "Casbin request: %v %v %v", UserID.(string), pageStr, timestampStr)

	// logic 层处理
	resp, err := logic.NewMessageLogic().GetMessage(c, UserID.(string), pageStr, timestampStr)

	// 响应
	if err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	} else {
		response.NewResponse(c).Success(resp)
	}

	return
}

// MarkReadMessage 标记已读, 更新 [用户-消息表] 的 read_at 字段
func MarkReadMessage(c *gin.Context) {
	// 解析请求参数
	ctx := zlog.GetCtxFromGin(c)
	req, err := types.BindReq[types.MarkReadMessageReq](c)
	if err != nil {
		return
	}
	zlog.CtxInfof(ctx, "Casbin request: %v", req)

	// logic 层处理
	resp, err := logic.NewMessageLogic().MarkReadMessage(c, req)

	// 响应
	if err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	} else {
		response.NewResponse(c).Success(resp)
	}

	return
}

// SendMessage 一键发送消息，SetMessage 和 JoinMessage 合并运用
func SendMessage(c *gin.Context) {
	// 解析请求参数
	ctx := zlog.GetCtxFromGin(c)
	req, err := types.BindReq[types.SendMessageReq](c)
	if err != nil {
		return
	}
	UserID, _ := c.Get(global.TOKEN_USER_ID)
	user_id, err := strconv.ParseInt(UserID.(string), 10, 64)
	if err != nil {
		return
	}
	zlog.CtxInfof(ctx, "Casbin request: %v", req)

	// logic 层处理
	resp, err := logic.NewMessageLogic().SendMessage(req, user_id)

	// 响应
	if err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	} else {
		response.NewResponse(c).Success(resp)
	}

	return
}
