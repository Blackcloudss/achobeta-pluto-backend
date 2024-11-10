package types

// 接收参数与返回参数层

type TestO1Req struct {
	UserID string `json:"user_id"`
}
type Test01Resp struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type PhoneReq struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
type PhoneResp struct {
	Atoken string `json:"atoken"`
	Rtoken string `json:"rtoken"`
}
