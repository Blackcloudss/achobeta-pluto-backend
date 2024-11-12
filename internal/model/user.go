package model

type User_Power struct {
	CommonModel
	member_id int64 `gorm:"column:member_id; type:bigint;   index;not null; comment:'成员ID'"`
	team_id   int64 `gorm:"column:team_id;   type:bigint;   index;not null; comment:'团队ID'"`
	level     int   `gorm:"column:level;     type:smallint; index;not null; default:1; comment:'权限等级'"`
}

func (t *User_Power) TableName() string {
	return "user_power"
}
