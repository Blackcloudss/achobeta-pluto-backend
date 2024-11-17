package repo

import (
	"gorm.io/gorm"
	"tgwp/global"
	"tgwp/internal/model"
	"tgwp/internal/types"
	"time"
)

type MessageRepo struct {
	DB *gorm.DB
}

func NewMessageRepo(db *gorm.DB) *MessageRepo {
	return &MessageRepo{DB: db}
}

// CheckMessageExists 检查消息是否存在
func (r MessageRepo) CheckMessageExist(message_id int64) bool {
	var count int64
	global.DB.Model(&model.Message{}).Where("id =?", message_id).Count(&count)
	if count > 0 {
		return true
	} else {
		return false
	}
}

// CheckUserMessageExists 检查用户消息是否存在
func (r MessageRepo) CheckUserMessageExist(user_message_id int64) bool {
	var count int64
	global.DB.Model(&model.UserMessage{}).Where("user_message_id =?", user_message_id).Count(&count)
	if count > 0 {
		return true
	} else {
		return false
	}
}

// CreateMessage 创建一条消息
func (r MessageRepo) CreateMessage(id int64, messageText string, messageType int) (message model.Message, err error) {
	message = model.Message{
		CommonModel: model.CommonModel{
			ID: id,
		},
		Content: messageText,
		Type:    messageType,
	}

	result := global.DB.Create(&message)
	err = result.Error

	return
}

// CreateUserMessage 连接一条用户消息
func (r MessageRepo) CreateUserMessage(id int64, message_id int64, user_id string) (user_message model.UserMessage, err error) {
	user_message = model.UserMessage{
		CommonModel: model.CommonModel{
			ID: id,
		},
		UserID:    user_id,
		MessageID: message_id,
		IsRead:    0,
	}

	result := global.DB.Create(&user_message)
	err = result.Error

	return
}

func (r MessageRepo) CheckUpdate(user_id string, timestamp int64) bool {
	FirstMessage := model.UserMessage{}
	global.DB.Model(&model.UserMessage{}).Where("user_id =?", user_id).Order("created_at desc").First(&FirstMessage)
	if FirstMessage.UpdatedAt.Unix() > timestamp {
		return true
	} else {
		return false
	}
}

// GetMessage 获取用户消息
func (r MessageRepo) GetMessage(user_id string, page int, pageSize int) (resp types.GetMessageResp, err error) {
	// 获取对应页面用户的消息id列表
	UserMessages := make([]model.UserMessage, pageSize)
	resp.Messages = make([]types.Message, pageSize)
	result := global.DB.Model(&model.UserMessage{}).Where("user_id =?", user_id).Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&UserMessages)
	err = result.Error

	// 获取总页数
	var total int64
	global.DB.Model(&model.UserMessage{}).Where("user_id =?", user_id).Count(&total)
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
			global.DB.Model(&model.Message{}).Where("id =?", UserMessages[i].MessageID).Find(&Message)
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
	result := global.DB.Model(&model.UserMessage{}).Where("id =?", UserMessageID).Update("is_read", 1)
	err = result.Error
	return
}
