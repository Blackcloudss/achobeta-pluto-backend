package repo

import (
	"fmt"
	"gorm.io/gorm"
	"tgwp/global"
	"tgwp/internal/model"
	"time"
)

const (
	SignTableName = "sign"
	OnlineTime    = "online_time"
	Issuer        = "issuer"
)

type SignRepo struct {
	DB *gorm.DB
}

func NewSignRepo(db *gorm.DB) *SignRepo {
	return &SignRepo{DB: db}
}

// InsertSign
//
//	@Description: 插入数据到sign表中
//	@receiver r
//	@param data
//	@return error
func (r SignRepo) InsertSign(data model.Sign) error {
	return global.DB.Table(SignTableName).
		Create(&data).Error
}

// CompareSign
//
//	@Description: 对比issuer是否有效
//	@receiver r
//	@param issuer
//	@return error
func (r SignRepo) CompareSign(issuer string) error {
	var data model.Sign
	return global.DB.Where(&model.Sign{Issuer: issuer}).First(&data).Error
}

// ReflashOnlineTime
//
//	@Description: 用于用户自动登录后，更新最新上线时间
//	@receiver r
//	@param issuer
func (r SignRepo) ReflashOnlineTime(issuer string) {
	global.DB.Table(SignTableName).Where(fmt.Sprintf("%s=?", Issuer), issuer).UpdateColumn(OnlineTime, time.Now())
}

//查找对应的Issuer并修改，自己退出登录
//根据LoginId修改issuer,被别人下线
