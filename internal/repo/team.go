package repo

import (
	"fmt"
	"gorm.io/gorm"
)

type TeamIdRepo struct {
	DB *gorm.DB
}

func NewTeamIdRepo(db *gorm.DB) *TeamIdRepo {
	return &TeamIdRepo{DB: db}
}

// 获取团队id
func (r TeamIdRepo) GetTeamId(userid int64) (fteamid int64, teamid []int64, err error) {
	err = r.DB.Table(TMSTableName).Select(TeamId).Where(fmt.Sprintf("%s = ? ", MemberId), userid).First(&fteamid).Error
	if err != nil {
		return
	}

	err = r.DB.Table(TeamTableName).Select(Id).Find(&teamid).Error
	if err != nil {
		return
	}
	return
}
