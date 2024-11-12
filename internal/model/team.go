package model

// 团队表
type Team struct {
	CommonModel
	TeamId string `gorm:"column:team_id; type:char(20); index; not null;  comment:'团队ID'"`
	Name   string `gorm:"column:name; type:varchar(50); index; not null;  comment:'分组/职位名字'"`
}
