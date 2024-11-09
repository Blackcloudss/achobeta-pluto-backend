package apig

import (
	"github.com/gin-gonic/gin"
	"tgwp/log/zlog"
)

// only for test
// 团队信息获取
func GetTeamStructure(c *gin.Context) {

}

func Test(c *gin.Context) {
	zlog.Infof("load - test")
	zlog.CtxInfof(c, "load ctx info - test")
	zlog.CtxWarnf(c, "load ctx warn - test")
	zlog.CtxErrorf(c, "load ctx error - test")
	c.JSON(200, gin.H{"msg": "success"})
	return
}
