package repo

import (
	"fmt"
	"gorm.io/gorm"
	"tgwp/global"
)

const (
	CasbinTableName = "casbin"
	Type            = "ptype"
	User            = "v0"
	Group           = "v1"
	Url             = "v2"
)

type CasbinRepo struct {
	DB *gorm.DB
}

func NewCasbinRepo(db *gorm.DB) *CasbinRepo {
	return &CasbinRepo{DB: db}
}

// 获取权限组
func (r CasbinRepo) Getcasbin(userid, teamid int64) ([]string, error) {
	var urls []string
	// 根据 UserId 查询用户对应的角色
	var roles []int64
	err := global.DB.Table(CasbinTableName).
		Select(Group). // 获取 g 规则中的 roleid
		Where(fmt.Sprintf("%s = 'g' AND %s = ?", Type, User), userid).
		Find(&roles).Error
	if err != nil {
		return nil, err
	}

	// 使用 roleid 和 teamid 查询拥有的 URL
	err = global.DB.Table(CasbinTableName).
		Select(Url). // 获取 p 规则中的 url
		Where(fmt.Sprintf("%s = 'p' AND %s IN ? AND %s = ?", Type, User, Group), roles, teamid).
		Find(&urls).Error
	if err != nil {
		return nil, err
	}

	return urls, nil
}
