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
		zlog.CtxErrorf(ctx, "BindReq failed: %v", err)
		err = response.ErrResp(err, response.INTERNAL_ERROR)
		return
	}
	zlog.CtxInfof(ctx, "SetMessage request: %v", req)

	// logic 层处理
	resp, err := logic.NewMessageLogic().SetMessage(req)

	// 响应
	response.Response(c, resp, err)
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
	TempUserID, ok := c.Get(global.TOKEN_USER_ID)
	if !ok {
		zlog.CtxErrorf(ctx, "Get token user id failed")
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	}
	UserID, ok := TempUserID.(int64)
	if !ok {
		zlog.CtxErrorf(ctx, "Token user id convert to int64 failed")
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	}

	zlog.CtxInfof(ctx, "JoinMessage request: %v", req)

	// logic 层处理
	resp, err := logic.NewMessageLogic().JoinMessage(req, UserID)

	// 响应
	response.Response(c, resp, err)

	return
}

// GetMessage 获取消息, 从 [消息表] 获取
func GetMessage(c *gin.Context) {
	//fmt.Println(util.GenToken(util.TokenData{Userid: 114514, Class: "atoken", Issuer: "", Time: time.Hour * 24 * 365}))

	// 解析请求参数
	ctx := zlog.GetCtxFromGin(c)

	req, err := types.BindReq[types.GetMessageReq](c)
	if err != nil {
		return
	}

	//获取token中的用户id
	TempUserID, ok := c.Get(global.TOKEN_USER_ID)
	if !ok {
		zlog.CtxErrorf(ctx, "Get token user id failed")
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	}
	UserID, ok := TempUserID.(int64)
	if !ok {
		zlog.CtxErrorf(ctx, "Token user id convert to int64 failed")
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	}

	zlog.CtxInfof(ctx, "GetMessage request: %v %v %v", UserID, req.Page, req.Timestamp)

	// logic 层处理
	resp, err := logic.NewMessageLogic().GetMessage(UserID, req)

	// 响应
	response.Response(c, resp, err)

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
	zlog.CtxInfof(ctx, "MarkReadMessage request: %v", req)

	// logic 层处理
	resp, err := logic.NewMessageLogic().MarkReadMessage(req)

	// 响应
	response.Response(c, resp, err)

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
	TempUserID, ok := c.Get(global.TOKEN_USER_ID)
	if !ok {
		zlog.CtxErrorf(ctx, "Get token user id failed")
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	}
	UserID, ok := TempUserID.(int64)
	if !ok {
		zlog.CtxErrorf(ctx, "Token user id convert to int64 failed")
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	}

	zlog.CtxInfof(ctx, "SendMessage request: %v", req)

	// logic 层处理
	resp, err := logic.NewMessageLogic().SendMessage(req, UserID)

	// 响应
	response.Response(c, resp, err)

	return
}
