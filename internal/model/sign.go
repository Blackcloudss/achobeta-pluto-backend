package model

import (
	"gorm.io/gorm"
	"time"
)

type Sign struct {
	LoginId   string `gorm:"column:login_id;type:char(19);primary;comment:'登录id'"`
	Issuer    string `gorm:"column:issuer;type:char(19);not null;comment:'签发标识'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (t *Sign) TableName() string {
	return "sign"
}
