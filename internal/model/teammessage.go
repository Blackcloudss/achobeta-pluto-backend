package model

import (
	"gorm.io/gorm"
	"time"
)

// 团队架构表
type team_structure struct {
	CommonModel
	TeamId     string `json:"team_id" gorm:"column:team_id;type:char(20);index;not null;comment:'团队ID'"`
	MyselfId   string `json:"myself_id" gorm:"column:myself_id;type:char(20);index;not null;comment:'当前节点ID'"`
	FatherId   string `json:"father_id" gorm:"column:father_id;type:char(20);index;not null;comment:'父节点ID'"`
	StructName string `json:"struct_name" gorm:"column:struct_name;type:varchar(50);index;not null;comment:'分组/职位名字'"`
}

type casbin_rule struct {
	ID uint `gorm:"primarykey"`

	Ptype string `gorm:"type:varchar(100);comment:'权限类型'"`
	V0    string `gorm:"type:char(20);index;comment:'用户ID'"`
	V1    string `gorm:"type:char(20);index;comment:'团队ID'"`
	V2    string `gorm:"type:varchar(100);index;comment:'用户的请求URL'"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
