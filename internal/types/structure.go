package types

// 查询团队架构(入参)
type Team_StructReq struct {
	UserId int64 `json:"userid"  binding:"required"`
	TeamId int64 `json:"team_id" binding:"required"`
}

// 单个团队架构记录
type TeamStructure struct {
	TeamId     int64  `json:"team_id"`        // 团队 ID
	MyselfId   int64  `json:"myself_id"`      // 当前节点 ID
	FatherId   int64  `json:"father_id"`      // 父节点 ID
	StructName string `json:"structure_name"` // 架构名称（职位名称等）
	IsDeleted  bool   `json:"is_deleted"`     //查看是否被删除
}

// 查询团队架构（出参）
type TeamStructResq struct {
	Code           int             `json:"code"`
	Message        string          `json:"msg"`
	TeamStructures []TeamStructure `json:"team_structures"` // 团队架构记录的数组
}

// 保存已经更改好的团队架构信息(入参）
type PostTeamNodeReq struct {
	UserId         int64           `json:"userid" binding:"required"`
	TeamId         int64           `json:"team_id" binding:"required"`
	TeamStructures []TeamStructure `json:"team_structures"` // 团队架构记录的数组
}

// 保存已经更改好的团队架构信息(出参）
type PostTeamNodeResq struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// 把团队架构信息变回 设定好初始值的团队架构(入参）
type DeleteTeamNodeReq struct {
	UserId int64 `json:"userid"`
	TeamId int64 `json:"team_id"`
}

// 把团队架构信息变回 设定好初始值的团队架构(出参）
type DeleteTeamNodeResq struct {
	Code           int             `json:"code"`
	Message        string          `json:"msg"`
	TeamStructures []TeamStructure `json:"team_structures"` // 团队架构记录的数组
}
