package logic

import (
	"tgwp/global"
	"tgwp/internal/repo"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
	"tgwp/util/snowflake"
)

type MessageLogic struct {
}

func NewMessageLogic() *MessageLogic {
	return &MessageLogic{}
}

// SetMessage
//
//	@Description: 设置初始消息
//	@receiver l
//	@param req
//	@return resp
//	@return err
func (l *MessageLogic) SetMessage(req types.SetMessageReq) (resp types.SetMessageResp, err error) {
	// 生成雪花ID生成器
	node, err := snowflake.NewNode(global.DEFAULT_NODE_ID)
	if err != nil {
		zlog.Errorf("create snowflake node error:%v", err)
		err = response.ErrResp(err, response.SNOWFLAKE_ID_GENERATE_ERROR)
		return
	}
	// 生成雪花ID
	id := node.Generate().Int64()

	message, err := repo.NewMessageRepo(global.DB).CreateMessage(id, req.Content, req.Type)
	if err != nil {
		zlog.Errorf("create message error:%v", err)
		err = response.ErrResp(err, response.DATABASE_ERROR)
		return
	} else {
		zlog.Infof("create message success, id:%d , content:\"%s\", type:%d", message.ID, message.Content, message.Type)
	}

	resp.MessageID = id

	return
}

// JoinMessage
//
//	@Description: 连接消息(和SetMessage配合)
//	@receiver l
//	@param req
//	@param UserID
//	@return resp
//	@return err
func (l *MessageLogic) JoinMessage(req types.JoinMessageReq, UserID int64) (resp types.JoinMessageResp, err error) {
	// 生成雪花ID生成器
	node, err := snowflake.NewNode(global.DEFAULT_NODE_ID)
	if err != nil {
		zlog.Errorf("create snowflake node error:%v", err)
		err = response.ErrResp(err, response.SNOWFLAKE_ID_GENERATE_ERROR)
		return
	}

	// 生成雪花ID
	id := node.Generate().Int64()

	// 更新数据库
	user_message, err := repo.NewMessageRepo(global.DB).CreateUserMessage(id, req.MessageID, UserID)
	if err != nil {
		zlog.Errorf("create message error:%v", err)
		err = response.ErrResp(err, response.DATABASE_ERROR)
		return
	} else {
		zlog.Infof("create message success, user_id:%d, message_id:%d", user_message.UserID, user_message.MessageID)
	}

	resp.UserMessageID = id

	return
}

// GetMessage
//
//	@Description: 获取消息
//	@receiver l
//	@param UserID
//	@param req
//	@return resp
//	@return err
func (l *MessageLogic) GetMessage(UserID int64, req types.GetMessageReq) (resp types.GetMessageResp, err error) {
	resp.IsUpdated = repo.NewMessageRepo(global.DB).CheckUpdate(UserID, req.Timestamp)
	if resp.IsUpdated {
		resp, err = repo.NewMessageRepo(global.DB).GetMessage(UserID, req.Page, 5)
		if err != nil {
			zlog.Errorf("get message error:%v", err)
			err = response.ErrResp(err, response.DATABASE_ERROR)
			return
		}
	}
	zlog.Infof("get message success")
	return
}

// MarkReadMessage
//
//	@Description: 标记已读
//	@receiver l
//	@param req
//	@return resp
//	@return err
func (l *MessageLogic) MarkReadMessage(req types.MarkReadMessageReq) (resp types.MarkReadMessageResp, err error) {

	// 更新数据库
	err = repo.NewMessageRepo(global.DB).MarkReadMessage(req.UserMessageID)

	if err != nil {
		zlog.Errorf("mark read error:%v", err)
		err = response.ErrResp(err, response.DATABASE_ERROR)
		return
	} else {
		zlog.Infof("mark read success, user_message_id:%d", req.UserMessageID)
	}

	return
}

// SendMessage
//
//	@Description: 发送消息
//	@receiver l
//	@param req
//	@param UserID
//	@return resp
//	@return err
func (l *MessageLogic) SendMessage(req types.SendMessageReq, UserID int64) (resp types.SendMessageResp, err error) {
	// 使用 SetMessage 方法
	respSet, err := l.SetMessage(types.SetMessageReq{
		Content: req.Content,
		Type:    req.Type,
	})
	if err != nil {
		return
	}
	// 使用 JoinMessage 方法
	respJoin, err := l.JoinMessage(types.JoinMessageReq{
		MessageID: respSet.MessageID,
	}, UserID)
	if err != nil {
		return
	}

	resp.MessageID = respSet.MessageID
	resp.UserMessageID = respJoin.UserMessageID
	return
}
