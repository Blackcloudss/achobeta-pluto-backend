package api

import (
	"github.com/gin-gonic/gin"
	"tgwp/global"
	"tgwp/internal/repo"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
	"tgwp/util"
)

// CheckAutoLogin
//
//	@Description: 验证是否可以自动登录
//	@param c
func CheckAutoLogin(c *gin.Context) {
	ctx := zlog.GetCtxFromGin(c)
	req, err := types.BindReq[types.TokenReq](c)
	if err != nil {
		return
	}
	zlog.CtxInfof(ctx, "CheckAutoLogin request: %v", req)
	data, err := util.IdentifyToken(ctx, req.Token)
	if err != nil {
		//对应token无效，直接让他返回
		response.NewResponse(c).Error(response.TOKEN_IS_EXPIRED)
		return
	}
	//判断其是否为rtoken
	if data.Class != global.AUTH_ENUMS_RTOKEN {
		response.NewResponse(c).Error(response.PARAM_TYPE_ERROR)
		return
	}
	err = repo.NewSignRepo(global.DB).CompareSign(data.Issuer)
	if err != nil {
		//表明找不到issuer相等的，即rtoken是无效的
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
		return
	}
	//到这里时，issuer一定有效，且唯一
	repo.NewSignRepo(global.DB).ReflashOnlineTime(data.Issuer)
	response.NewResponse(c).Success(response.SUCCESS)
}
