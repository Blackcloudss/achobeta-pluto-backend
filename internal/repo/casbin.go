package repo

import (
	"tgwp/global"
)

const CasbinTableName = "casbin"

type CasbinRepo struct {
}

func NewCasbinRepo() *CasbinRepo {
	return &CasbinRepo{}
}

// 获取权限组
func (r CasbinRepo) Getcasbin(userid, teamid int64) ([]string, error) {
	var urls []string
	// 根据 UserId 查询用户对应的角色
	var roles []int64
	err := global.DB.Table(CasbinTableName).
		Select("v1"). // 获取 g 规则中的 roleid
		Where("ptype = 'g' AND v0 = ?", userid).
		Find(&roles).Error
	if err != nil {
		return nil, err
	}

	// 使用 roleid 和 teamid 查询拥有的 URL
	err = global.DB.Table(CasbinTableName).
		Select("v2"). // 获取 p 规则中的 url
		Where("ptype = 'p' AND v0 IN ? AND v1 = ?", roles, teamid).
		Find(&urls).Error
	if err != nil {
		return nil, err
	}

	return urls, nil
}
