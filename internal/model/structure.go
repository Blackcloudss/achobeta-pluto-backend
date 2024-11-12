package model

// 团队架构表
type Team_Structure struct {
	CommonModel
	TeamId     string `gorm:"column:team_id; type:char(20); index;not null; comment:'团队ID'"`
	MyselfId   string `gorm:"column:myself_id; type:char(20); index;not null; comment:'当前节点ID'"`
	FatherId   string `gorm:"column:father_id; type:char(20); index; not null; comment:'父节点ID'"`
	StructName string `gorm:"column:struct_name; type:varchar(50); index; not null; comment:'分组/职位名字'"`
}
