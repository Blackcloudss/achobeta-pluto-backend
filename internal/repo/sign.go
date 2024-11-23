package repo

import (
	"fmt"
	"gorm.io/gorm"
	"tgwp/internal/model"
	"tgwp/internal/types"
	"time"
)

const (
	SignTableName = "sign"
	OnlineTime    = "online_time"
	Issuer        = "issuer"
	Phone         = "phone"
	UserId        = "user_id"
	LoginId       = "login_id"
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
	return r.DB.Table(SignTableName).
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
	return r.DB.Where(&model.Sign{Issuer: issuer}).Take(&data).Error
}

// ReflashOnlineTime
//
//	@Description: 用于用户自动登录后，更新最新上线时间
//	@receiver r
//	@param issuer
func (r SignRepo) ReflashOnlineTime(issuer string) {
	r.DB.Table(SignTableName).Where(fmt.Sprintf("%s=?", Issuer), issuer).UpdateColumn(OnlineTime, time.Now())
}

// CheckUserId
//
//	@Description: 根据手机号查找用户是否已经有过userid，确保userid唯一
//	@receiver r
//	@param phone
func (r SignRepo) CheckUserId(phone string) (int64, error) {
	//建立一个临时结构体
	var Temp struct {
		UserId int64 `gorm:"column:user_id"` // 假设你的数据库列名是 user_id
	}
	err := r.DB.Table(SignTableName).Select(UserId).
		Where(fmt.Sprintf("%s=?", Phone), phone).
		Take(&Temp).Error
	if err != nil {
		return 0, err
	}
	// 返回检索到的 user_id
	return Temp.UserId, nil
}

// DeleteSignByIssuer
//
//	@Description: 查找对应的Issuer并删除，自己退出登录
//	@receiver r
//	@param issuer
//	@return err
func (r SignRepo) DeleteSignByIssuer(issuer string) (err error) {
	var Temp model.Sign
	err = r.DB.Table(SignTableName).Where(fmt.Sprintf("%s=?", Issuer), issuer).Delete(&Temp).Error
	return
}

// DeleteSignByLoginId
//
//	@Description: 根据LoginId删除信息,被别人下线
//	@receiver r
//	@param login_id
//	@return err
func (r SignRepo) DeleteSignByLoginId(login_id int64) (err error) {
	var Temp model.Sign
	err = r.DB.Table(SignTableName).Where(fmt.Sprintf("%s=?", LoginId), login_id).Delete(&Temp).Error
	return
}

// ShowDevices
//
//	@Description: 展示常用设备
//	@receiver r
//	@param user_id
//	@return err
func (r SignRepo) ShowDevices(req types.DevicesReq) (resq types.DevicesResp, err error) {
	fmt.Println(req.PageNumber, req.LineNumber)
	offset := (req.PageNumber - 1) * req.LineNumber
	r.DB.Model(&model.Sign{}).
		Where(fmt.Sprintf("%s=?", UserId), req.UserId).
		Count(&resq.Total)
	if resq.Total == 0 {
		return
	}
	err = r.DB.Model(&model.Sign{}).
		Where(fmt.Sprintf("%s=?", UserId), req.UserId).
		Offset(offset).
		Limit(req.LineNumber).
		Find(&resq.Devices).Error
	return
}
