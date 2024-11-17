package initalize

import (
	"tgwp/global"
	"tgwp/internal/model"
)

func MigrateTables() {
	// 自动迁移 casbin 表，确保表结构存在
	global.DB.AutoMigrate(&model.Casbin{})

	// 自动迁移 team 表，确保表结构存在
	global.DB.AutoMigrate(&model.Team{})

	// 自动迁移 member 表，确保表结构存在
	global.DB.AutoMigrate(&model.Member{})

	// 自动迁移 like_status 表，确保表结构存在
	global.DB.AutoMigrate(&model.Like_Status{})

	// 自动迁移 structure 表，确保表结构存在
	global.DB.AutoMigrate(&model.Structure{})

	// 自动迁移 team_member_structure 表，确保表结构存在
	global.DB.AutoMigrate(&model.Team_Member_Structure{})

	// 自动迁移 user_power 表，确保表结构存在
	global.DB.AutoMigrate(&model.User_Power{})
}
