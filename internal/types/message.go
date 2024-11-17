package types

// SetMessageReq 设置消息请求（入参）
type SetMessageReq struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

// SetMessageResp 设置消息请求（出参）
type SetMessageResp struct {
	MessageID int64 `json:"message_id"` // 消息ID
}

// JoinMessageReq 连接消息请求（入参）
type JoinMessageReq struct {
	Atoken    string `json:"atoken"`
	MessageID int64  `json:"message_id"`
}

// JoinMessageResp 连接消息请求（出参）
type JoinMessageResp struct {
}

// GetMessageReq 获取消息请求（入参）
type GetMessageReq struct {
	Atoken    string `json:"atoken"`
	Page      int    `json:"page"`
	Timestamp int64  `json:"timestamp"`
}

// Message 消息
type Message struct {
	UserMessageID int64  `json:"user_message_id"`
	MessageID     int64  `json:"message_id"`
	Type          int    `json:"type"`
	Content       string `json:"content"`
	ReceivedAt    int64  `json:"received_at"`
	IsRead        int    `json:"is_read"`
}

// GetMessageResp 获取消息请求（出参）
type GetMessageResp struct {
	IsUpdated  bool      `json:"is_updated"`
	TotalPages int       `json:"total_pages"`
	Messages   []Message `json:"messages"`
}

// MarkReadMessageReq 标记已读消息请求（入参）
type MarkReadMessageReq struct {
	UserMessageID int64 `json:"user_message_id"`
}

// MarkReadMessageResp 标记已读消息请求（出参）
type MarkReadMessageResp struct {
}
