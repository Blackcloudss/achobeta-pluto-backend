package types

// 查询团队架构(入参)
type Team_StructReq struct {
	UserId string `json:"userid"`
	TeamId string `json:"team_id"`
}

// 单个团队架构记录
type TeamStructure struct {
	TeamId     string `json:"team_id"`        // 团队 ID
	MyselfId   string `json:"myself_id"`      // 当前节点 ID
	FatherId   string `json:"father_id"`      // 父节点 ID
	StructName string `json:"structure_name"` // 架构名称（职位名称等）
}

// 查询团队架构（出参）
type TeamStructResq struct {
	Code           int             `json:"code"`
	Message        string          `json:"msg"`
	TeamStructures []TeamStructure `json:"team_structures"` // 团队架构记录的数组
}

// 新职位的ID发给前端(入参）
type GetTeamNodeReq struct {
	UserId   string `json:"userid"`
	TeamId   string `json:"team_id"`
	FatherId string `json:"father_id"`
}

// 新职位的ID发给前端(出参）
type GetTeamNodeResq struct {
	Code     int    `json:"code"`
	Message  string `json:"msg"`
	MyselfId string `json:"myself_id"` // 新节点 ID
}

// 保存已经更改好的团队架构信息(入参）
type PostTeamNodeReq struct {
	UserId         string          `json:"userid"`
	TeamId         string          `json:"team_id"`
	TeamStructures []TeamStructure `json:"team_structures"` // 团队架构记录的数组
}

// 保存已经更改好的团队架构信息(出参）
type PostTeamNodeResq struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// 把团队架构信息变回 设定好初始值的团队架构(入参）
type DeleteTeamNodeReq struct {
	UserId string `json:"userid"`
	TeamId string `json:"team_id"`
}

// 把团队架构信息变回 设定好初始值的团队架构(出参）
type DeleteTeamNodeResq struct {
	Code           int             `json:"code"`
	Message        string          `json:"msg"`
	TeamStructures []TeamStructure `json:"team_structures"` // 团队架构记录的数组
}
