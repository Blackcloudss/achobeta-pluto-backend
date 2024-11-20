package repo

import (
	"gorm.io/gorm"
	"tgwp/internal/model"
	"tgwp/internal/types"
)

type TeamIdRepo struct {
	DB *gorm.DB
}

const c_teamid = "team_id"
const c_teamname = "name"

func NewTeamIdRepo(db *gorm.DB) *TeamIdRepo {
	return &TeamIdRepo{DB: db}
}

func (r TeamIdRepo) GetTeamId(userid int64) (first_team types.Team, team []types.Team, err error) {

	err = r.DB.Model(&model.Team_Member_Structure{}).
		Select(c_teamid, c_teamname).
		Joins("JOIN team on team.id = team_member_structure.team_id").
		Where(&model.Team_Member_Structure{
			MemberId: userid,
		}).
		First(&first_team).
		Error
	if err != nil {
		return
	}

	err = r.DB.Model(&model.Team{}).
		Joins("JOIN team on team.id = team_member_structure.team_id").
		Select(C_Id, c_teamname).
		Find(&team).
		Error
	if err != nil {
		return
	}
	return
}
