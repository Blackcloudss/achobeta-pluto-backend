package model

type User_Power struct {
	CommonModel
	member_id string `gorm:"column:member_id; type:char(20); index;not null; comment:'成员ID'"`
	team_id   string `gorm:"column:team_id;   type:char(20); index;not null; comment:'团队ID'"`
	level     int    `gorm:"column:level;     type:smallint; default:1; index;not null; comment:'权限等级'"`
}
