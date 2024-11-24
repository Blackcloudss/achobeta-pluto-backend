package types

// 获得权限组（入参）
type RuleReq struct {
	UserId int64 `form:"user_id"` // 仅为了测试使用，之后删除
	TeamId int64 `form:"team_id"`
}

type Team struct {
	TeamId   int64  `gorm:"column:id" json:"id"`
	TeamName string `gorm:"column:name" json:"name"`
}

// 获得权限组（出参）
type RuleResp struct {
	Url           []string `json:"urls"`         // 包含权限 URL 的数组
	Level         int      `json:"level"`        // 权限等级
	FirstTeamID   int64    `json:"first_teamid"` // 第一个团队 ID
	FirstTeamName string   `json:"first_team_name"`
	Team          []Team   `json:"teams"` // 团队 ID 数组
}

// 表单权限验证
type ParamsRuleCheck struct {
	UserId int64 `form:"user_id"` // 测试时使用
	TeamId int64 `form:"team_id"`
}

// Json权限验证
type JsonRuleCheck struct {
	UserId int64 `json:"user_id"` // 测试时使用
	TeamId int64 `json:"team_id"`
}

// Uri权限验证
type UriRuleCheck struct {
	UserId int64 `uri:"user_id"` // 测试时使用
	TeamId int64 `uri:"team_id"`
}
