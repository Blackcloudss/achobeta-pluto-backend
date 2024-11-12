package repo

import (
	"tgwp/global"
)

const TeamTableName = "team"

type TeamIdRepo struct {
}

func NewTeamIdRepo() *TeamIdRepo {
	return &TeamIdRepo{}
}

// 获取团队id
func GetTeamId() (fteamid int64, teamid []int64, err error) {
	global.DB.Table(TeamTableName).First(&fteamid)
	global.DB.Table(TeamTableName).Find(&teamid)
	return
}
