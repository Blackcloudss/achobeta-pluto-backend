package repo

import (
	"gorm.io/gorm"
	"tgwp/internal/model"
	"tgwp/internal/types"
	"tgwp/log/zlog"
	"time"
)

type MessageRepo struct {
	db *gorm.DB
}

func NewMessageRepo(db *gorm.DB) *MessageRepo {
	return &MessageRepo{db: db}
}

// CheckMessageExists
//
//	@Description: 检查消息是否存在
//	@receiver r
//	@param MessageID
//	@return isExists
//	@return err
func (r MessageRepo) CheckMessageExists(MessageID int64) (isExists bool, err error) {
	err = r.db.Model(&model.Message{}).Where("id = ?", MessageID).Find(&isExists).Error
	return
}

// CheckUserMessageExists
//
//	@Description: 检查用户消息是否存在
//	@receiver r
//	@param UserMessageID
//	@return isExists
//	@return err
func (r MessageRepo) CheckUserMessageExists(UserMessageID int64) (isExists bool, err error) {
	err = r.db.Model(&model.UserMessage{}).Where("id = ?", UserMessageID).Find(&isExists).Error
	return
}

// CreateMessage 创建一条消息
func (r MessageRepo) CreateMessage(messageText string, messageType int) (message model.Message, err error) {
	message = model.Message{
		CommonModel: model.CommonModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Content: messageText,
		Type:    messageType,
	}

	result := r.db.Model(&model.Message{}).Create(&message)
	err = result.Error

	// 需要返回hook生成的雪花ID
	var ok bool
	message.ID, ok = result.Statement.ReflectValue.FieldByName("ID").Interface().(int64)
	if !ok {
		zlog.Errorf("message.ID is not int64")
		return
	}
	return
}

// CreateUserMessage 连接一条用户消息
func (r MessageRepo) CreateUserMessage(message_id int64, user_id int64) (user_message model.UserMessage, err error) {
	user_message = model.UserMessage{
		CommonModel: model.CommonModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UserID:    user_id,
		MessageID: message_id,
		IsRead:    0,
	}

	result := r.db.Create(&user_message)
	err = result.Error

	// 需要返回hook生成的雪花ID
	var ok bool
	user_message.ID, ok = result.Statement.ReflectValue.FieldByName("ID").Interface().(int64)
	if !ok {
		zlog.Errorf("message.ID is not int64")
		return
	}
	return
}

func (r MessageRepo) CheckUpdate(user_id int64, timestamp int64) bool {
	FirstMessage := model.UserMessage{}
	var cnt int64
	r.db.Model(&model.UserMessage{}).Where("user_id =?", user_id).Count(&cnt)
	if cnt == 0 { // 没有消息，要返回已更新
		return true
	}

	r.db.Model(&model.UserMessage{}).Where("user_id =?", user_id).Order("created_at desc").First(&FirstMessage)
	if FirstMessage.UpdatedAt.Unix() > timestamp {
		return true
	} else {
		return false
	}
}

// GetMessage 获取用户消息
func (r MessageRepo) GetMessage(user_id int64, page int, pageSize int) (resp types.GetMessageResp, err error) {
	// 获取对应页面用户的消息id列表
	UserMessages := make([]model.UserMessage, pageSize)
	resp.Messages = make([]types.Message, pageSize)
	result := r.db.Model(&model.UserMessage{}).Where("user_id =?", user_id).Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&UserMessages)
	err = result.Error

	// 获取总页数
	var total int64
	r.db.Model(&model.UserMessage{}).Where("user_id =?", user_id).Count(&total)
	resp.TotalPages = int(total / int64(pageSize))
	if total%int64(pageSize) != 0 {
		resp.TotalPages += 1
	}

	//用消息id获取对应的消息内容
	for i := 0; i < pageSize; i++ {
		if i >= len(UserMessages) {
			// 没有消息
			resp.Messages[i] = types.Message{
				UserMessageID: 0,
				MessageID:     0,
				Content:       "",
				Type:          0,
				ReceivedAt:    0,
				IsRead:        0,
			}
		} else {
			// 获取消息内容
			Message := model.Message{}
			r.db.Model(&model.Message{}).Where("id =?", UserMessages[i].MessageID).Find(&Message)
			resp.Messages[i] = types.Message{
				UserMessageID: UserMessages[i].ID,
				MessageID:     Message.ID,
				Content:       Message.Content,
				Type:          Message.Type,
				ReceivedAt:    time.Now().Unix(),
				IsRead:        UserMessages[i].IsRead,
			}
		}
	}

	return
}

func (r MessageRepo) MarkReadMessage(UserMessageID int64) (err error) {
	result := r.db.Model(&model.UserMessage{}).Where("id =?", UserMessageID).Update("is_read", 1)
	err = result.Error
	return
}
