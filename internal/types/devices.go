package types

type Device struct {
	Id         string `json:"id"`
	Ip         string `json:"ip"`
	OnlineTime string `json:"online_time"`
	DeviceName string `json:"device_name"`
}
type DevicesReq struct {
	PageNumber int   `form:"page_number"`
	LineNumber int   `form:"line_number"`
	UserId     int64 `json:"user_id"`
}
type DevicesResp struct {
	Total   int64    `json:"total"`
	Devices []Device `json:"devices"`
}
type RemoveDeviceReq struct {
	Id int64 `json:"id"`
}
type ModifyDeviceNameReq struct {
	Id         int64  `json:"id"`
	DeviceName string `json:"device_name"`
}
