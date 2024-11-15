package repo

import (
	"gorm.io/gorm"
	"tgwp/global"
	"tgwp/internal/model"
)

const (
	SignTableName = "sign"
)

type SignRepo struct {
	DB *gorm.DB
}

func NewSignRepo(db *gorm.DB) *SignRepo {
	return &SignRepo{DB: db}
}

// 插入数据
func (r SignRepo) InsertSign(login_id, issuer string) error {
	data := model.Sign{
		LoginId: login_id,
		Issuer:  issuer,
	}
	return global.DB.Table(SignTableName).
		Create(&data).Error
}

// 对比issuer是否有效
func (r SignRepo) CompareSign(issuer string) error {
	var data model.Sign
	return global.DB.Where(&model.Sign{Issuer: issuer}).First(&data).Error
}

//查找对应的Issuer并修改，自己退出登录
//根据LoginId修改issuer,被别人下线
