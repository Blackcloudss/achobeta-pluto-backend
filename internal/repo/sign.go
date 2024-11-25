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
	Issuer        = "issuer"
	UserId        = "user_id"
	LoginId       = "id"
	DeviceName    = "device_name"
	CreatedAt     = "created_at"
)

type SignRepo struct {
	DB *gorm.DB
}

func NewSignRepo(db *gorm.DB) *SignRepo {
	return &SignRepo{DB: db}
}

// InsertSign
//
//	@Description: 插入数据到sign表中,同时返回雪花id作为login_id
//	@receiver r
//	@param data
//	@return error
func (r SignRepo) InsertSign(data model.Sign) (id int64, err error) {
	err = r.DB.Table(SignTableName).
		Create(&data).Error
	return data.ID, err
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
	r.DB.Table(SignTableName).
		Where(fmt.Sprintf("%s=?", Issuer), issuer).
		Updates(model.Sign{OnlineTime: time.Now()})
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
func (r SignRepo) ShowDevices(req types.DevicesReq) (resp types.DevicesResp, err error) {
	offset := (req.PageNumber - 1) * req.LineNumber
	// 计算30天前的日期
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
	//由于我们每次进行的都是删除操作，只要一条数据已经创建超过30天，那么这个rtoken必定失效了
	r.DB.Model(&model.Sign{}).
		Where(fmt.Sprintf("%s=?", UserId), req.UserId).
		Where(fmt.Sprintf("%s>?", CreatedAt), thirtyDaysAgo).
		Count(&resp.Total)
	if resp.Total == 0 {
		return
	}
	err = r.DB.Model(&model.Sign{}).
		Where(fmt.Sprintf("%s=?", UserId), req.UserId).
		Where(fmt.Sprintf("%s>?", CreatedAt), thirtyDaysAgo).
		Offset(offset).
		Limit(req.LineNumber).
		Find(&resp.Devices).Error
	return
}

// ModifyDeviceName
//
//	@Description: 根据设备的login_id修改设备名称
//	@receiver r
//	@param login_id
//	@param device_name
//	@return err
func (r SignRepo) ModifyDeviceName(req types.ModifyDeviceNameReq) (err error) {
	err = r.DB.Table(SignTableName).
		Where(fmt.Sprintf("%s=?", LoginId), req.LoginId).
		UpdateColumn(DeviceName, req.DeviceName).Error
	return
}
