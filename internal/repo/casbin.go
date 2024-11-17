package repo

import (
	"fmt"
	"gorm.io/gorm"
	"tgwp/internal/model"
)

type CasbinRepo struct {
	DB *gorm.DB
}

func NewCasbinRepo(db *gorm.DB) *CasbinRepo {
	return &CasbinRepo{DB: db}
}

// GetCasbin
//
//	@Description:
//	@receiver r
//	@param userid
//	@param teamid
//	@return []string
//	@return error
//
// 获取权限组
func (r CasbinRepo) GetCasbin(userid, teamid int64) ([]string, error) {
	// 根据 UserId 查询用户对应的角色
	var roles []int64
	err := r.DB.Model(&model.Casbin{}).
		Select(RoleOrUrl). // 获取 g 规则中的 roleid
		Where(&model.Casbin{
			Ptype: "g",
			V0:    userid,
			V1:    teamid,
		}).
		Find(&roles).Error
	if err != nil {
		return nil, err
	}

	// 使用 roleid 和 teamid 查询拥有的 URL
	var urls []string
	err = r.DB.Model(&model.Casbin{}).
		Select(RoleOrUrl). // 获取 p 规则中的 url
		Where(fmt.Sprintf("%s = 'p' AND %s in (?) AND %s = ?", C_Type, C_User, C_Team), roles, teamid).
		Find(&urls).Error
	if err != nil {
		return nil, err
	}

	return urls, nil
}
