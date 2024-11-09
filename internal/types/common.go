package types

import (
	"github.com/gin-gonic/gin"
	"tgwp/internal/response"
)

func BindReq[T any](c *gin.Context) (req T, err error) {
	if err = c.ShouldBind(&req); err != nil {
		response.NewResponse(c).Error(response.PARAM_NOT_COMPLETE)
	}
	return req, err
}
