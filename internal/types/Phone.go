package types

type PhoneReq struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
type PhoneResp struct {
	Atoken string `json:"atoken"`
	Rtoken string `json:"rtoken"`
}
