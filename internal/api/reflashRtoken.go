package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/global"
	"tgwp/internal/logic"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
	"tgwp/util"
)

func ReflashRtoken(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)
	req, err := types.BindReq[types.TokenReq](c)
	if err != nil {
		return
	}
	//解析token是否有效，并取出上一次的值
	data, err := util.IdentifyToken(ctx, req.Token)
	if err != nil {
		zlog.CtxErrorf(ctx, "ReflashRtoken err:%v", err)
		response.NewResponse(c).Error(response.TOKEN_IS_EXPIRED)
		//对应token无效，直接让他返回
		return
	}
	//判断其是否为rtoken
	if data.Class != global.AUTH_ENUMS_RTOKEN {
		response.NewResponse(c).Error(response.PARAM_TYPE_ERROR)
		return
	}
	//生成新的token
	resp, err := logic.NewTokenLogic().GenRtoken(ctx, data)
	if err != nil {
		zlog.CtxErrorf(ctx, "ReflashRtoken err:%v", err)
		return
	}
	response.NewResponse(c).Success(resp)
}
