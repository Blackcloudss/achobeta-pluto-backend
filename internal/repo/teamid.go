package repo

import (
	"gorm.io/gorm"
	"tgwp/internal/model"
)

type TeamIdRepo struct {
	DB *gorm.DB
}

func NewTeamIdRepo(db *gorm.DB) *TeamIdRepo {
	return &TeamIdRepo{DB: db}
}

// GetTeamId
//
//	@Description:
//	@receiver r
//	@param userid
//	@return fteamid
//	@return teamid
//	@return err
//
// 获取团队id
func (r TeamIdRepo) GetTeamId(userid int64) (fteamid int64, teamid []int64, err error) {
	err = r.DB.Model(&model.Team_Member_Structure{}).
		Select(C_TeamId).
		Where(&model.Team_Member_Structure{
			MemberId: userid,
		}).
		First(&fteamid).Error
	if err != nil {
		return
	}

	err = r.DB.Model(&model.Team{}).
		Select(C_Id).
		Find(&teamid).Error
	if err != nil {
		return
	}
	return
}
