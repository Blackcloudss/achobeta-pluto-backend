package types

type SetMessageReq struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

type JoinMessageReq struct {
	Atoken    string `json:"atoken"`
	MessageID int64  `json:"message_id"`
}
