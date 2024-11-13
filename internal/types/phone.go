package types

type PhoneReq struct {
	Phone     string `json:"phone"`
	Code      string `json:"code"`
	AutoLogin bool   `json:"auto_login"`
}
type PhoneResp struct {
	Atoken    string `json:"atoken"`
	Rtoken    string `json:"rtoken"`
	ServiceId uint   `json:"service_id"`
	UserAgent string `json:"user_agent"`
	Ip        string `json:"ip"`
}
