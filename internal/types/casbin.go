package types

// 获得权限组（入参）
type RuleReq struct {
	UserId int64 `json:"user_id"  binding:"required"`
	TeamId int64 `json:"team_id"`
}

// 获得权限组（出参）
type RuleResp struct {
	Url         []string `json:"url"`          // 包含权限 URL 的数组
	Level       int      `json:"level"`        // 权限等级
	FirstTeamID int64    `json:"first_teamid"` // 第一个团队 ID
	TeamID      []int64  `json:"teamid"`       // 团队 ID 数组
}

// 权限验证
type RuleCheck struct {
	UserId int64 `json:"user_id"  binding:"required"`
	TeamId int64 `json:"team_id"`
}
