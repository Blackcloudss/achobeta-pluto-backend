package repo

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"tgwp/internal/model"
	"tgwp/log/zlog"
	"tgwp/util"
	"time"
)

type CasbinRepo struct {
	DB *gorm.DB
}

func NewCasbinRepo(db *gorm.DB) *CasbinRepo {
	return &CasbinRepo{DB: db}
}

const Nothing = 0

// GetCasbin
// @Description: 获取权限组
// @receiver r
// @param userid 用户ID
// @param teamid 团队ID
// @return level 用户权限等级
// @return urls 用户拥有的URL列表
// @return error
func (r CasbinRepo) GetCasbin(userid, teamid int64) (int, []string, error) {
	// 根据 UserId 查询用户对应的角色
	var rolesWithLevels []struct {
		Role  int64
		Level int
	}

	// 修复 JOIN 表别名冲突
	err := r.DB.Model(&model.Casbin{}).
		Joins("JOIN user_power AS up1 ON up1.member_id = casbin.v0").
		Joins("JOIN user_power AS up2 ON up2.team_id = casbin.v1").
		Select("casbin.v2 AS role, up2.level").
		Where("casbin.ptype = ? AND casbin.v0 = ? AND casbin.v1 = ?", "g", userid, teamid).
		Find(&rolesWithLevels).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Nothing, nil, nil
		}
		zlog.Errorf("查询用户对应的用户组失败：%v", err)
		return Nothing, nil, err
	}

	if len(rolesWithLevels) == 0 {
		return Nothing, nil, nil
	}

	level := rolesWithLevels[0].Level
	var roles []int64
	for _, r := range rolesWithLevels {
		roles = append(roles, r.Role)
	}

	// 使用角色ID和团队ID查询拥有的URL
	var urls []string
	err = r.DB.Model(&model.Casbin{}).
		Select("casbin.v2").
		Where("casbin.ptype = ? AND casbin.v0 IN ? AND casbin.v1 = ?", "p", roles, teamid).
		Find(&urls).Error
	if err != nil {
		zlog.Errorf("查询管理员拥有的 URLs 失败：%v", err)
		return Nothing, nil, err
	}

	return level, urls, nil
}

// CheckUserPermission
//
//	@Description: 检查用户权限
//	@receiver r
//	@param url
//	@param userId
//	@param teamId
//	@return bool
//	@return error
func (r CasbinRepo) CheckUserPermission(url string, userId, teamId int64) (bool, error) {
	defer util.RecordTime(time.Now())()

	var roles []string
	err := r.DB.Model(&model.Casbin{}).
		Select(RoleOrUrl). // 获取 g 规则中的 roleid
		Where(&model.Casbin{
			Ptype: "g",
			V0:    userId,
			V1:    teamId,
		}).
		Find(&roles).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		zlog.Errorf("查询用户对应的用户组失败：%v", err)
		return false, err
	}

	var managers []int64
	for _, role := range roles {
		//将string类型转化为 int64
		manager, err := strconv.ParseInt(role, 10, 64)
		if err != nil {
			//转换失败
			return false, err
		}
		managers = append(managers, manager)
	}

	var res string
	// 使用 roleid 和 teamid 查询拥有的 URL
	err = r.DB.Model(&model.Casbin{}).
		Select(RoleOrUrl). // 获取 p 规则中的 url
		Where(fmt.Sprintf("%s = 'p' AND %s in (?) AND %s = ? AND %s = ?", C_Type, C_User, C_Team, RoleOrUrl), managers, teamId, url).
		First(&res).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 查询不到记录，无权限
			return false, nil
		}
		zlog.Errorf("权限验证出错：%v", err)
		return false, err
	}

	zlog.Infof("权限验证成功: managers=%v, teamId=%d, url=%s", managers, teamId, url)

	// 查询成功，存在记录，有权限
	return true, nil
}
