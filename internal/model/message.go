package model

// Message [信息表]
type Message struct {
	CommonModel
	Content string `json:"content" ;gorm:"type:varchar(2048);null;comment:'消息具体内容'"`
	Type    int    `json:"type" ;gorm:"type:int;null;comment:'消息类型'"`
}

type UserMessage struct {
	CommonModel
	MessageID int64  `json:"message_id"`
	UserID    string `json:"user_id"`
	IsRead    int    `json:"is_read"`
}
