package types

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/response"
)

// BindReq 一个通用的方法，用来绑定请求参数，当请求参数绑定失败时自动返回错误 使用泛型来更加通用
func BindReq[T any](c *gin.Context) (req T, err error) {
	if err = c.ShouldBindJSON(&req); err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_VALID)
	}
	return req, err
}
