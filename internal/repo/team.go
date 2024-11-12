package repo

import (
	"tgwp/global"
	"tgwp/internal/model"
)

// 获取团队id
func GetTeamId() (fteamid string, teamid []string, err error) {
	global.DB.AutoMigrate(&model.Team{})
	global.DB.Table("team").First(&fteamid)
	global.DB.Table("team").Find(&teamid)
	return
}
