package model

// 团队架构表
type Structure struct {
	CommonModel
	TeamId     int64  `gorm:"column:team_id;     type:bigint; index; not null; comment:'团队ID'"`
	MyselfId   int64  `gorm:"column:myself_id;   type:bigint; index; comment:'当前节点ID'"`
	FatherId   int64  `gorm:"column:father_id;   type:bigint; index; comment:'父节点ID'"`
	StructName string `gorm:"column:struct_name; type:varchar(50); index; not null; comment:'分组/职位名字'"`
	IsDeleted  bool   `gorm:"column:is_deleted; type:unsigned tinyint; index;not null; comment:'删除情况' "`
}

func (t *Structure) TableName() string {
	return "structure"
}
