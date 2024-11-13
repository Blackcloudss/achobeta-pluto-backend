package repo

import (
	"gorm.io/gorm"
	"tgwp/global"
)

const (
	TeamTableName = "team"
	TeamId        = "id"
)

type TeamIdRepo struct {
	DB *gorm.DB
}

func NewTeamIdRepo(db *gorm.DB) *TeamIdRepo {
	return &TeamIdRepo{DB: db}
}

// 获取团队id
func (r TeamIdRepo) GetTeamId() (fteamid int64, teamid []int64, err error) {
	err = global.DB.Table(TeamTableName).Select(TeamId).First(&fteamid).Error
	if err != nil {
		return
	}

	err = global.DB.Table(TeamTableName).Select(TeamId).Find(&teamid).Error
	if err != nil {
		return
	}
	return
}
