package model

// Message [信息表]
type Message struct {
	CommonModel
	Content string `json:"content" ;gorm:"type:varchar(2048);null;comment:'消息具体内容'"`
	Type    int    `json:"type" ;gorm:"type:int;null;comment:'消息类型'"`
}

func (t *Message) TableName() string {
	return "messages"
}

// UserMessage [用户消息表]
type UserMessage struct {
	CommonModel
	MessageID int64 `json:"message_id"`
	UserID    int64 `json:"user_id"`
	IsRead    int   `json:"is_read"`

	Message Message `gorm:"foreignKey:MessageID;references:ID;"`
}

func (t *UserMessage) TableName() string {
	return "user_messages"
}
