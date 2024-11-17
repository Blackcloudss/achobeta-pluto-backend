package repo

import (
	"gorm.io/gorm"
	"tgwp/internal/model"
	"tgwp/util"
	"time"
)

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
	var res string
	err := r.DB.Model(&model.Casbin{}).
		Select(RoleOrUrl).
		Where(&model.Casbin{
			V0: userId,
			V1: teamId,
			V2: url,
		}).
		First(&res).
		Error
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
