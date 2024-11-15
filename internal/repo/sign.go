package repo

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"tgwp/global"
)

const (
	SignTableName = "sign"
	UserId        = "user_id"
	UserAgent     = "user_agent"
	IP            = "ip"
	Issuer        = "issuer"
)

type SignRepo struct {
	DB *gorm.DB
}
type SignData struct {
	IP        string
	UserAgent string
	UserId    string
	Issuer    string
}

func NewSignRepo(db *gorm.DB) *SignRepo {
	return &SignRepo{DB: db}
}

// 插入一条数据，登陆时，若存在数据，改掉对应的issuer
func (r SignRepo) UpdateSign(data SignData) error {
	result := global.DB.Table(SignTableName).
		Where(fmt.Sprintf("%s=? And %s=? And %s=?", UserId, UserAgent, IP), data.UserId, data.UserAgent, data.IP)
	if result.Error != nil {
		// 如果发生错误，并且错误不是因为记录未找到（GORM返回gorm.ErrRecordNotFound表示记录未找到）
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error // 返回查找时的错误
		}
		// 如果没有找到记录，则创建新记录
		return global.DB.Table(SignTableName).Create(&data).Error // 创建新用户并返回错误（如果有的话）
	}
	// 如果找到了记录，则更新Issuer字段
	return global.DB.Table(SignTableName).Update(fmt.Sprintf("%s=?", Issuer), data.Issuer).Error
}

//根据UserAgent,ip,UserId查找对应的Issuer，自己退出登录
//根据LoginId修改issuer,被别人下线
