package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tgwp/global"
	"tgwp/internal/model"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
	"tgwp/util"
	"tgwp/util/snowflake"
	"time"
)

// SetMessage 设置消息, 存入 [消息表]
func SetMessage(c *gin.Context) {
	// 解析请求参数
	req, err := types.BindReq[types.SetMessageReq](c)
	if err != nil {
		return
	}

	// 生成雪花ID生成器
	node, err := snowflake.NewNode(global.DEFAULT_NODE_ID)
	if err != nil {
		zlog.Errorf("create snowflake node error:%v", err)
		return
	}
	// 生成雪花ID
	ID := node.Generate().Int64()

	message := model.Message{
		ID:      ID,
		Content: req.Content,
		Type:    req.Type,
	}

	result := global.DB.Create(&message)
	if result.Error != nil {
		zlog.Errorf("create message error:%v", result.Error)
	} else {
		zlog.Infof("create message success, id:%d , content:\"%s\", type:%d", message.ID, message.Content, message.Type)
	}

	response.NewResponse(c).Success(
		struct {
			MessageId int64 `json:"message_id"`
		}{
			MessageId: message.ID,
		})
	return
}

// JoinMessage 连接用户与消息, 存入 [用户-消息表]
func JoinMessage(c *gin.Context) {
	// 解析请求参数
	req, err := types.BindReq[types.JoinMessageReq](c)
	if err != nil {
		return
	}
	fmt.Println(util.GenToken(util.TokenData{Userid: "114514", Class: "1", Issuer: "1", Time: time.Hour * 24 * 365}))

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
	// 生成雪花ID
	ID := node.Generate().Int64()

	user_message := model.UserMessage{
		ID:        ID,
		UserID:    data.Userid,
		MessageID: req.MessageID,
		IsRead:    0,
	}

	// 判断信息是否存在
	var count int64
	global.DB.Table("messages").Where("id =?", req.MessageID).Count(&count)
	if count == 0 {
		zlog.Warnf("message not exist, message_id:%d", req.MessageID)
		return
	}

	// 更新数据库
	result := global.DB.Create(&user_message)
	if result.Error != nil {
		zlog.Errorf("create message error:%v", result.Error)
	} else {
		zlog.Infof("create message success, user_id:%d, message_id:%d", user_message.UserID, user_message.MessageID)
	}

	return
}
