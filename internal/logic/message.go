package logic

import (
	"github.com/gin-gonic/gin"
	"tgwp/global"
	"tgwp/internal/repo"
	"tgwp/internal/types"
	"tgwp/log/zlog"
	"tgwp/util"
	"tgwp/util/snowflake"
)

type SetMessageLogic struct {
}

type JoinMessageLogic struct {
}

func NewSetMessageLogic() *SetMessageLogic {
	return &SetMessageLogic{}
}

func NewJoinMessageLogic() *JoinMessageLogic {
	return &JoinMessageLogic{}
}

func (l *SetMessageLogic) SetMessage(c *gin.Context, req types.SetMessageReq) (resp types.SetMessageResp, err error) {
	// 生成雪花ID生成器
	node, err := snowflake.NewNode(global.DEFAULT_NODE_ID)
	if err != nil {
		zlog.Errorf("create snowflake node error:%v", err)
		return
	}
	// 生成雪花ID
	id := node.Generate().Int64()

	message, err := repo.NewSetMessageRepo(global.DB).CreateMessage(id, req.Content, req.Type)
	if err != nil {
		zlog.Errorf("create message error:%v", err)
	} else {
		zlog.Infof("create message success, id:%d , content:\"%s\", type:%d", message.ID, message.Content, message.Type)
	}

	resp.MessageID = id

	return
}

func (l *JoinMessageLogic) JoinMessage(c *gin.Context, req types.JoinMessageReq) (resp types.JoinMessageResp, err error) {
	// 解析token
	data, err := util.ParseToken(req.Atoken)
	if err != nil {
		zlog.Warnf("parse token error:%v", err)
		return
	}

	// 生成雪花ID生成器
	node, err := snowflake.NewNode(global.DEFAULT_NODE_ID)
	if err != nil {
		zlog.Errorf("create snowflake node error:%v", err)
		return
	}

	// 判断信息是否存在
	var count int64
	global.DB.Table("messages").Where("id =?", req.MessageID).Count(&count)
	if count == 0 {
		zlog.Warnf("message not exist, message_id:%d", req.MessageID)
		return
	}

	// 生成雪花ID
	id := node.Generate().Int64()

	user_message, err := repo.NewJoinMessageRepo(global.DB).CreateUserMessage(id, req.MessageID, data.Userid)
	// 更新数据库
	if err != nil {
		zlog.Errorf("create message error:%v", err)
	} else {
		zlog.Infof("create message success, user_id:%d, message_id:%d", user_message.UserID, user_message.MessageID)
	}

	return
}
