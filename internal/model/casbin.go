package model

import (
	"gorm.io/gorm"
	"time"
)

type casbin_rule struct {
	ID        uint   `gorm:"primarykey"`
	Ptype     string `gorm:"type:varchar(100);comment:'权限类型'"`
	V0        string `gorm:"type:char(20);index;comment:'用户ID'"`
	V1        string `gorm:"type:char(20);index;comment:'团队ID'"`
	V2        string `gorm:"type:varchar(100);index;comment:'用户的请求URL'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
