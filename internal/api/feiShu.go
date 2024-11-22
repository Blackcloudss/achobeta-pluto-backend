package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/global"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

func GetFeiShuList(c *gin.Context) {
	// 解析请求参数
	ctx := zlog.GetCtxFromGin(c)
	TempUserID, _ := c.Get(global.TOKEN_USER_ID)
	UserID := TempUserID.(int64)

	req, err := types.BindReq[types.GetFeiShuListReq](c)
	if err != nil {
		return
	}

	zlog.CtxInfof(ctx, "GetFeiShuList requset %v", UserID)

	// logic 层处理
	resp, err := logic.NewFeiShuLogic().GetFeiShuList(ctx, UserID, req.ForceUpdate)

	// 响应
	if err != nil {
		response.Response(c, resp, err)
		return
	} else {
		response.NewResponse(c).Success(resp)
	}
	return
}
