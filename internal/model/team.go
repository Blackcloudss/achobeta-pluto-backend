package model

// 团队表
type Team struct {
	CommonModel
	TeamId int64  `gorm:"column:team_id; type:bigint; index; not null;  comment:'团队ID'"`
	Name   string `gorm:"column:name; type:varchar(50); index; not null;  comment:'分组/职位名字'"`
}

func (t *Team) TableName() string {
	return "team"
}
