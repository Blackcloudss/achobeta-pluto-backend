package model

import (
	"gorm.io/gorm"
	"time"
)

// CommonModel 每张表都有的四个东西，最好不要用 gorm.model（虽然他们一模一样）
type CommonModel struct {
	ID        int64 `gorm:"primary;column:id;type:bigint"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
