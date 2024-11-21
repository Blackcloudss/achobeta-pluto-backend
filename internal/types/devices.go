package types

type Device struct {
	LoginId    string `json:"login_id"`
	Ip         string `json:"ip"`
	OnlineTime string `json:"online_time"`
	DeviceName string `json:"device_name"`
}
type DevicesReq struct {
	PageNumber int   `json:"page_number"`
	LineNumber int   `json:"line_number"`
	UserId     int64 `json:"user_id"`
}
type DevicesResp struct {
	Token   string   `json:"token"`
	Total   int64    `json:"total"`
	Devices []Device `json:"devices"`
}
type RemoveDeviceReq struct {
	LoginId string `json:"login_id"`
}
