package types

// 获得权限组（入参）
type RuleReq struct {
	UserId string `json:"userid"`
	TeamId string `json:"team_id"`
}

// 获得权限组（出参）
type RuleResq struct {
	Code        int      `json:"code"`
	Message     string   `json:"message"`
	Data        []string `json:"data"`        // 包含权限 URL 的数组
	Level       int      `json:"level"`       // 权限等级
	FirstTeamID string   `json:"firstteamid"` // 第一个团队 ID
	TeamID      []string `json:"teamid"`      // 团队 ID 数组
}
