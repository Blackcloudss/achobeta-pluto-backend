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
