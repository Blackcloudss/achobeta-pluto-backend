package repo

import (
	"tgwp/global"
	"tgwp/internal/model"
	"tgwp/internal/types"
)

// 获取权限组
func Getcasbin(req types.RuleReq) ([]string, error) {
	var urls []string

	// 自动迁移 Casbin 模型，确保表结构存在
	global.DB.AutoMigrate(&model.Casbin{})

	// 根据 UserId 查询用户对应的角色
	var roles []string
	err := global.DB.Table("casbin_rule").
		Select("v1"). // 获取 g 规则中的 roleid
		Where("ptype = 'g' AND v0 = ?", req.UserId).
		Find(&roles).Error
	if err != nil {
		return nil, err
	}

	// 使用 roleid 和 teamid 查询拥有的 URL
	err = global.DB.Table("casbin_rule").
		Select("v2"). // 获取 p 规则中的 url
		Where("ptype = 'p' AND v0 IN ? AND v1 = ?", roles, req.TeamId).
		Find(&urls).Error
	if err != nil {
		return nil, err
	}

	return urls, nil
}
