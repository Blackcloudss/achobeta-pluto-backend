package model

import (
	"gorm.io/gorm"
	"time"
)

// CommonModel 每张表都有的四个东西，最好不要用 gorm.model（虽然他们一模一样）
type CommonModel struct {
	id        int64 `gorm:"primarykey;type:bigint"`
	createdat time.Time
	updatedat time.Time
	deletedat gorm.DeletedAt `gorm:"index"`
}
