package model

// 团队表
type Team struct {
	CommonModel
	Name string `gorm:"column:name; type:varchar(50); unique; not null;  comment:'团队名字'"`
}

func (t *Team) TableName() string {
	return "team"
}
