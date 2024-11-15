package repo

import (
	"fmt"
	"gorm.io/gorm"
	"tgwp/util"
	"time"
)

type CasbinRepo struct {
	DB *gorm.DB
}

func NewCasbinRepo(db *gorm.DB) *CasbinRepo {
	return &CasbinRepo{DB: db}
}

// 获取权限组
func (r CasbinRepo) GetCasbin(userid, teamid int64) ([]string, error) {
	// 根据 UserId 查询用户对应的角色
	var roles int64
	err := r.DB.Table(CasbinTableName).
		Select(RoleOrUrl). // 获取 g 规则中的 roleid
		Where(fmt.Sprintf("%s = 'g' AND %s = ? AND %s = ?", C_Type, C_User, C_Team), userid, teamid).
		First(&roles).Error
	if err != nil {
		return nil, err
	}

	// 使用 roleid 和 teamid 查询拥有的 URL
	var urls []string
	err = r.DB.Table(CasbinTableName).
		Select(RoleOrUrl). // 获取 p 规则中的 url
		Where(fmt.Sprintf("%s = 'p' AND %s = ? AND %s = ?", C_Type, C_User, C_Team), roles, teamid).
		Find(&urls).Error
	if err != nil {
		return nil, err
	}

	return urls, nil
}

// 获取权限级别
type LevelRepo struct {
	DB *gorm.DB
}

func NewLevelRepo(db *gorm.DB) *LevelRepo {
	return &LevelRepo{DB: db}
}

func (r LevelRepo) GetLevel(userid, teamid int64) (int, error) {
	var level int
	err := r.DB.Table(UPTableName).
		Select(C_Level).
		Where(fmt.Sprintf("%s = ? AND %s = ?", C_MemberId, C_TeamId), userid, teamid).
		First(&level).Error
	if err != nil {
		return 0, err
	}

	return level, nil
}

type PermissionRepo struct {
	DB *gorm.DB
}

func NewPermissionRepo(db *gorm.DB) *PermissionRepo {
	return &PermissionRepo{DB: db}
}

// 查询权限
// CheckUserPermissions 检查用户权限
func (r PermissionRepo) CheckUserPermission(url string, userId, teamId int64) error {
	defer util.RecordTime(time.Now())()
	err := r.DB.Table(CasbinTableName).
		Select(RoleOrUrl).
		Where(fmt.Sprintf("%s = ? AND %s = ? AND %s = ?", C_User, C_Team, RoleOrUrl), userId, teamId, url).
		Error
	return err
}
