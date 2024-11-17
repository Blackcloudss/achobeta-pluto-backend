package repo

import (
	"gorm.io/gorm"
	"tgwp/global"
	"tgwp/internal/model"
	"time"
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

type CommonData struct {
	LoginId    string    `json:"login_id"`
	Issuer     string    `json:"issuer"`
	UserId     string    `json:"user_id"`
	IP         string    `json:"ip"`
	UserAgent  string    `json:"user_agent"`
	OnlineTime time.Time `json:"online_time"`
}

// InsertSign
//
//	@Description: 插入数据到sign表中
//	@receiver r
//	@param data
//	@return error
func (r SignRepo) InsertSign(data CommonData) error {
	temp := model.Sign{
		LoginId:    data.LoginId,
		Issuer:     data.Issuer,
		UserId:     data.UserId,
		IP:         data.IP,
		UserAgent:  data.UserAgent,
		OnlineTime: data.OnlineTime,
	}
	return global.DB.Table(SignTableName).
		Create(&temp).Error
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

//查找对应的Issuer并修改，自己退出登录
//根据LoginId修改issuer,被别人下线
