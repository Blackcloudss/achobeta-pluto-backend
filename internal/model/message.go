package model

// Message [信息表]
type Message struct {
	CommonModel
	ID      int64  `json:"id"`
	Content string `json:"content"`
	Type    int    `json:"type"`
}

type UserMessage struct {
	CommonModel
	ID        int64  `json:"id"`
	MessageID int64  `json:"message_id"`
	UserID    string `json:"user_id"`
	IsRead    int    `json:"is_read"`
}
