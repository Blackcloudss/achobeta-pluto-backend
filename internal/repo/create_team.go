package repo

import (
	"gorm.io/gorm"
	"tgwp/global"
	"tgwp/internal/model"
	"tgwp/internal/types"
	"time"
)

type CreateTeamRepo struct {
	DB *gorm.DB
}

func NewCreateTeamRepo(db *gorm.DB) *CreateTeamRepo {
	return &CreateTeamRepo{
		DB: db,
	}
}

// CreateTeam
//
//	@Description: 新增团队 ，初始化团队架构
//	@receiver r
//	@param TeamName
//	@return types.CreateTeamResp
//	@return error
func (r CreateTeamRepo) CreateTeam(TeamName string) (types.CreateTeamResp, error) {

	//创建新团队
	err := r.DB.Model(&model.Team{}).
		Create(&model.Team{Name: TeamName}).
		Error
	if err != nil {
		return types.CreateTeamResp{}, err
	}

	//找到新创建的团队ID
	var TeamId int64
	err = r.DB.Model(&model.Team{}).
		Select(C_Id).
		Where(&model.Team{Name: TeamName}).
		First(&TeamId).
		Error
	if err != nil {
		return types.CreateTeamResp{}, err
	}

	//初始化团队架构
	err = r.DB.Model(&model.Structure{}).
		Create(&model.Structure{
			CommonModel: model.CommonModel{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			FatherId: global.ROOT_ID,
			NodeName: TeamName,
			TeamId:   TeamId,
		}).
		Error
	if err != nil {
		return types.CreateTeamResp{}, err
	}

	//初始化团队权限组
	Rules := []*model.Casbin{}

	// 普通管理员
	for _, url := range global.NORMAL_ADMIN_URLS {
		rule := &model.Casbin{
			CommonModel: model.CommonModel{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Ptype: "p",
			V0:    global.NORMAL_ADMINISTRATOR,
			V1:    TeamId,
			V2:    url,
		}
		Rules = append(Rules, rule)
	}
	// 超级管理员
	for _, url := range global.SUPER_ADMIN_URLS {
		rule := &model.Casbin{
			CommonModel: model.CommonModel{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Ptype: "p",
			V0:    global.SUPERL_ADMINISTRATOR, // 超级管理员
			V1:    TeamId,
			V2:    url,
		}
		Rules = append(Rules, rule)
	}

	for _, Rule := range Rules {
		err = r.DB.Model(&model.Casbin{}).
			Create(Rule).
			Error
	}

	if err != nil {
		return types.CreateTeamResp{}, err
	}
	return types.CreateTeamResp{}, nil
}
