package repo

import (
	"gorm.io/gorm"
	"tgwp/global"
	"tgwp/internal/model"
)

type SetMessageRepo struct {
	DB *gorm.DB
}

type JoinMessageRepo struct {
	DB *gorm.DB
}

func NewSetMessageRepo(db *gorm.DB) *SetMessageRepo {
	return &SetMessageRepo{DB: db}
}

func NewJoinMessageRepo(db *gorm.DB) *JoinMessageRepo {
	return &JoinMessageRepo{DB: db}
}

// CreateMessage 创建一条消息
func (r SetMessageRepo) CreateMessage(id int64, messageText string, messageType int) (message model.Message, err error) {
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
func (r JoinMessageRepo) CreateUserMessage(id int64, message_id int64, user_id string) (user_message model.UserMessage, err error) {
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
