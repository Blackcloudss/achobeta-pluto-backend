package repo

import (
	"gorm.io/gorm"
	"tgwp/global"
	"tgwp/internal/model"
	"tgwp/internal/types"
	"time"
)

type PostTeamRepo struct {
	DB *gorm.DB
}

func NewPostTeamRepo(db *gorm.DB) *PostTeamRepo {
	return &PostTeamRepo{
		DB: db,
	}
}

// PostTeam
//
//	@Description:
//	@receiver r
//	@param TeamName
//	@return types.PostTeamResp
//	@return error
func (r PostTeamRepo) PostTeam(TeamName string) (types.PostTeamResp, error) {

	//创建新团队
	err := r.DB.Model(&model.Team{}).
		Create(&model.Team{Name: TeamName}).
		Error
	if err != nil {
		return types.PostTeamResp{}, err
	}

	//找到新创建的团队ID
	var TeamId int64
	err = r.DB.Model(&model.Team{}).
		Select(C_Id).
		Where(&model.Team{Name: TeamName}).
		First(&TeamId).
		Error
	if err != nil {
		return types.PostTeamResp{}, err
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
		return types.PostTeamResp{}, err
	}

	//初始化团队权限组
	Rules := []*model.Casbin{
		//普通管理员所拥有的
		{
			CommonModel: model.CommonModel{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Ptype: "p",
			V0:    global.NORMAL_ADMINISTRATOR,
			V1:    TeamId,
			V2:    "/api/team/memberlist/delete",
		},
		{
			CommonModel: model.CommonModel{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Ptype: "p",
			V0:    global.NORMAL_ADMINISTRATOR,
			V1:    TeamId,
			V2:    "/api/team/memberlist/put",
		},
		{
			CommonModel: model.CommonModel{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Ptype: "p",
			V0:    global.NORMAL_ADMINISTRATOR,
			V1:    TeamId,
			V2:    "/api/team/members/save",
		},
		{
			CommonModel: model.CommonModel{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Ptype: "p",
			V0:    global.NORMAL_ADMINISTRATOR,
			V1:    TeamId,
			V2:    "/api/team/structure/collection",
		},
		//超级管理员所拥有的
		{
			CommonModel: model.CommonModel{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Ptype: "p",
			V0:    global.SUPERL_ADMINISTRATOR,
			V1:    TeamId,
			V2:    "/api/team/structure/change",
		},
		{
			CommonModel: model.CommonModel{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Ptype: "p",
			V0:    global.SUPERL_ADMINISTRATOR,
			V1:    TeamId,
			V2:    "/api/team/structure/add",
		},
	}

	for _, Rule := range Rules {
		err = r.DB.Model(&model.Casbin{}).
			Create(Rule).
			Error
	}

	if err != nil {
		return types.PostTeamResp{}, err
	}
	return types.PostTeamResp{}, nil
}
