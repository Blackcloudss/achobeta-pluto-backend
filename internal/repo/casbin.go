package repo

import (
	"fmt"
	"gorm.io/gorm"
	"tgwp/internal/model"
	"tgwp/util"
	"time"
)

type CasbinRepo struct {
	DB *gorm.DB
}

const Nothing = 0

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
func (r CasbinRepo) GetCasbin(userid, teamid int64) (int, []string, error) {
	// 根据 UserId 查询用户对应的角色
	var Power []struct {
		role  int64
		level int
	}

	err := r.DB.Model(&model.Casbin{}).
		Joins("JOIN user_power ON user_power.member_id = casbin.v0").
		Joins("JOIN user_power ON user_power.team_id = casbin.v1").
		Select(RoleOrUrl, C_Level). // 获取 g 规则中的 roleid
		Where(&model.Casbin{
			Ptype: "g",
			V0:    userid,
			V1:    teamid,
		}).
		Find(&Power).Error
	if err != nil {
		return Nothing, nil, err
	}
	Level := Power[0].level

	var roles []int64
	for _, power := range Power {
		roles = append(roles, power.role)
	}

	// 使用 roleid 和 teamid 查询拥有的 URL
	var urls []string
	err = r.DB.Model(&model.Casbin{}).
		Select(RoleOrUrl). // 获取 p 规则中的 url
		Where(fmt.Sprintf("%s = 'p' AND %s in (?) AND %s = ?", C_Type, C_User, C_Team), roles, teamid).
		Find(&urls).Error
	if err != nil {
		return Nothing, nil, err
	}

	return Level, urls, nil
}

// 查询用户权限
type PermissionRepo struct {
	DB *gorm.DB
}

func NewPermissionRepo(db *gorm.DB) *PermissionRepo {
	return &PermissionRepo{DB: db}
}

// CheckUserPermission
//
//	@Description:
//	@receiver r
//	@param url
//	@param userId
//	@param teamId
//	@return bool
//	@return error
//
// 查询权限
// CheckUserPermissions 检查用户权限
func (r PermissionRepo) CheckUserPermission(url string, userId, teamId int64) (bool, error) {
	defer util.RecordTime(time.Now())()

	var roles []int64
	err := r.DB.Model(&model.Casbin{}).
		Select(RoleOrUrl). // 获取 g 规则中的 roleid
		Where(&model.Casbin{
			Ptype: "g",
			V0:    userId,
			V1:    teamId,
		}).
		Find(&roles).Error
	if err != nil {
		return false, err
	}

	var res string
	// 使用 roleid 和 teamid 查询拥有的 URL
	err = r.DB.Model(&model.Casbin{}).
		Select(RoleOrUrl). // 获取 p 规则中的 url
		Where(fmt.Sprintf("%s = 'p' AND %s in (?) AND %s = ? AND %s = ?", C_Type, C_User, C_Team), roles, teamId, url).
		First(&res).Error

	//查询出错
	if err != nil {
		return false, err
	}
	//查询成功
	if res == "" {
		//没有记录：无权限
		return false, err
	} else {
		//有记录：有权限
		return true, err
	}
}
