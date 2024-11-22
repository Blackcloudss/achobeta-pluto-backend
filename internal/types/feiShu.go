package types

// GetFeiShuListReq 获取飞书多维表格请求（入参）
type GetFeiShuListReq struct {
	ForceUpdate bool `form:"force-update"`
}

// GetFeiShuListResp 获取飞书多维表格请求（出参）
type GetFeiShuListResp struct {
	TotalTaskCount       int `json:"total_task_count"`
	UnFinishedTaskCount  int `json:"unfinished_task_count"`
	WillOverdueTaskCount int `json:"will_overdue_task_count"`
	OverdueTaskCount     int `json:"overdue_task_count"`
}
