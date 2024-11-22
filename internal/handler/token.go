package handler

import (
	"github.com/gin-gonic/gin"
	"tgwp/global"
)

func GetUserId(c *gin.Context) int64 {
	if data, exists := c.Get(global.TOKEN_USER_ID); exists {
		return data.(int64)
	}
	return 0
}
