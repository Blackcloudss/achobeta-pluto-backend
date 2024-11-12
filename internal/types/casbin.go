package types

// 获得权限组（入参）
type RuleReq struct {
	UserId int64 `json:"user_id"  binding:"required"`
	TeamId int64 `json:"team_id"`
}

// 获得权限组（出参）
type RuleResq struct {
	Code        int      `json:"code"`
	Message     string   `json:"message"`
	Data        []string `json:"data"`        // 包含权限 URL 的数组
	Level       int      `json:"level"`       // 权限等级
	FirstTeamID int64    `json:"firstteamid"` // 第一个团队 ID
	TeamID      []int64  `json:"teamid"`      // 团队 ID 数组
}
