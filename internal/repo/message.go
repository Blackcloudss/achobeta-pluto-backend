package repo

import (
	"gorm.io/gorm"
	"tgwp/global"
	"tgwp/internal/model"
)

type MessageRepo struct {
	DB *gorm.DB
}

func NewMessageRepo(db *gorm.DB) *MessageRepo {
	return &MessageRepo{DB: db}
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
