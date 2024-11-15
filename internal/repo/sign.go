package repo

import "gorm.io/gorm"

type SignRepo struct {
	DB *gorm.DB
}

func NewSignRepo(db *gorm.DB) *SignRepo {
	return &SignRepo{DB: db}
}
