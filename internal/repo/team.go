package repo

import (
	"gorm.io/gorm"
	"tgwp/global"
)

const TeamTableName = "team"

type TeamIdRepo struct {
	DB *gorm.DB
}

func NewTeamIdRepo(db *gorm.DB) *TeamIdRepo {
	return &TeamIdRepo{DB: db}
}

// 获取团队id
func (r TeamIdRepo) GetTeamId() (fteamid int64, teamid []int64, err error) {
	global.DB.Table(TeamTableName).First(&fteamid)
	global.DB.Table(TeamTableName).Find(&teamid)
	return
}
