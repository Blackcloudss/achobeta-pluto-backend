package repo

import (
	"gorm.io/gorm"
	"strconv"
	"tgwp/global"
	"tgwp/internal/model"
	"tgwp/internal/types"
	"tgwp/util"
	"time"
)

const (
	init_likecount = 0
	team_id        = "id"
	init_teamname  = "未分配团队"
)

// 将 int64 转化成 string
var (
	NormalManger = strconv.FormatInt(global.NORMAL_ADMINISTRATOR, 10)
	SuperManger  = strconv.FormatInt(global.SUPERL_ADMINISTRATOR, 10)
)

type CreateMemberRepo struct {
	DB *gorm.DB
}

func NewCreateMemberRepo(db *gorm.DB) *CreateMemberRepo {
	return &CreateMemberRepo{
		DB: db,
	}
}

// CreateMember
//
//	@Description: 新增团队成员
//	@receiver r
//	@param req
//	@return error
func (r *CreateMemberRepo) CreateMember(req types.CreateMemberReq) error {
	defer util.RecordTime(time.Now())()
	err := r.DB.Model(&model.Member{}).
		Create(&model.Member{
			CommonModel: model.CommonModel{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:       req.Name,
			Sex:        req.Sex,
			CreateDate: req.CreateDate,
			IdCard:     req.IdCard,
			PhoneNum:   req.PhoneNum,
			Email:      req.Email,
			Grade:      req.Grade,
			Major:      req.Major,
			StudentID:  req.StudentID,
			Experience: req.Experience,
			Status:     req.Status,
			LikeCount:  init_likecount,
		}).
		Error
	if err != nil {
		return err
	}

	//获取新成员id
	var UserId int64

	err = r.DB.Model(&model.Member{}).
		Select(C_Id).
		Where(&model.Member{
			PhoneNum: req.PhoneNum,
		}).
		First(&UserId).
		Error
	if err != nil {
		return err
	}

	if req.MemberPositions == nil {
		//说明没有给他分配任何团队，那么把他放入“未分配团队”的团队中
		var InitTeamId int64

		err := r.DB.Model(&model.Team{}).
			Select(team_id).
			Where("name = ?", init_teamname).
			First(&InitTeamId).
			Error

		if err != nil {
			return err
		}

		if InitTeamId == 0 {
			//说明没有“未分配团队”的团队id不存在，重新创建
			_, err = NewCreateTeamRepo(global.DB).CreateTeam(init_teamname)
			if err != nil {
				return err
			}
		}
		//重新获取
		err = r.DB.Model(&model.Team{}).
			Select(team_id).
			Where("name = ?", init_teamname).
			First(&InitTeamId).
			Error
		if err != nil {
			return err
		}

		// 新成员id 和 未分配团队id 在 Team_Member_Structure 做一个简单关联
		err = r.DB.Model(&model.Team_Member_Structure{}).
			Create(&model.Team_Member_Structure{
				MemberId: UserId,
				TeamId:   InitTeamId,
			}).Error
		if err != nil {
			return err
		}
	}

	for _, member_position := range req.MemberPositions {
		for _, structure_node := range member_position.StructNodes {
			// 新成员id 和 团队id 对 Team_Member_Structure表 初始化
			// 多职位id关联
			err = r.DB.Model(&model.Team_Member_Structure{}).
				Create(&model.Team_Member_Structure{
					MemberId:    UserId,
					TeamId:      member_position.TeamId,
					StructureId: structure_node.ID,
				}).Error
			if err != nil {
				return err
			}
		}
		// 对 User_Power表 初始化
		err = r.DB.Model(&model.User_Power{}).
			Create(&model.User_Power{
				MemberId: UserId,
				TeamId:   member_position.TeamId,
				Level:    member_position.Level,
			}).Error
		if err != nil {
			return err
		}

		// 根据权限级别 对 Casbin表 初始化
		if member_position.Level >= 2 {
			err = r.DB.Model(&model.Casbin{}).
				Create(&model.Casbin{
					Ptype: "g",
					V0:    UserId,
					V1:    member_position.TeamId,
					V2:    NormalManger,
				}).Error
			if err != nil {
				return err
			}
		}
		if member_position.Level == 3 {
			err = r.DB.Model(&model.Casbin{}).
				Create(&model.Casbin{
					Ptype: "g",
					V0:    UserId,
					V1:    member_position.TeamId,
					V2:    SuperManger,
				}).Error
			if err != nil {
				return err
			}
		}

	}

	return nil
}

type DeleteMemberRepo struct {
	DB *gorm.DB
}

func NewDeleteMemberRepo(db *gorm.DB) *DeleteMemberRepo {
	return &DeleteMemberRepo{
		DB: db,
	}
}

// DeleteMember
//
//	@Description: 删除团队成员
//	@receiver r
//	@param MemberId
//	@param TeamId
//	@return error
func (r *DeleteMemberRepo) DeleteMember(MemberId, TeamId int64) error {
	defer util.RecordTime(time.Now())()

	//删除 用户 和 团队 关联的职位
	err := r.DB.Model(&model.Team_Member_Structure{}).
		Delete(&model.Team_Member_Structure{}).
		Where(&model.Team_Member_Structure{
			MemberId: MemberId,
			TeamId:   TeamId,
		}).Error
	if err != nil {
		return err
	}

	//删除 用户 和 团队 关联的权限级别
	err = r.DB.Model(&model.User_Power{}).
		Delete(&model.User_Power{}).
		Where(&model.User_Power{
			MemberId: MemberId,
			TeamId:   TeamId,
		}).Error
	if err != nil {
		return err
	}

	//删除 用户 和 团队 关联的 casbin表
	err = r.DB.Model(&model.Casbin{}).
		Delete(&model.Casbin{}).
		Where(&model.Casbin{
			Ptype: "g",
			V0:    MemberId,
			V1:    TeamId,
		}).Error
	if err != nil {
		return err
	}

	//删除 用户 和 团队 关联的 点赞表
	err = r.DB.
		Where("memberid_like = ? OR memberid_belike = ?", MemberId, MemberId).
		Delete(&model.Like_Status{}).
		Error
	if err != nil {
		return err
	}

	return nil
}

type PutMemberRepo struct {
	DB *gorm.DB
}

func NewPutMemberRepo(db *gorm.DB) *PutMemberRepo {
	return &PutMemberRepo{
		DB: db,
	}
}

// PutMember
//
//	@Description: 编辑团队成员
//	@receiver r
//	@param req
//	@return error
func (r *PutMemberRepo) PutMember(req types.PutTeamMemberReq) error {
	defer util.RecordTime(time.Now())()

	//更改基本信息
	err := r.DB.Model(&model.Member{}).Updates(&model.Member{
		Name:       req.Name,
		Sex:        req.Sex,
		CreateDate: req.CreateDate,
		IdCard:     req.IdCard,
		PhoneNum:   req.PhoneNum,
		Email:      req.Email,
		Grade:      req.Grade,
		Major:      req.Major,
		StudentID:  req.StudentID,
		Experience: req.Experience,
		Status:     req.Status,
	}).Where(&model.Member{
		CommonModel: model.CommonModel{
			ID: req.ID,
		},
		PhoneNum: req.PhoneNum,
	}).Error
	if err != nil {
		return err
	}

	/// 删除旧职位信息
	err = r.DB.Where("member_id = ?", req.ID).
		Delete(&model.Team_Member_Structure{}).
		Error
	if err != nil {
		//
		return err
	}

	// 插入新的职位信息
	for _, position := range req.MemberPosition {
		for _, node := range position.StructNodes {
			newStructure := model.Team_Member_Structure{
				TeamId:      position.TeamId,
				MemberId:    req.ID,
				StructureId: node.ID,
			}
			if err := r.DB.Create(&newStructure).Error; err != nil {
				//
				return err
			}
		}

		//更新 user_power权限表
		err = r.DB.Model(&model.User_Power{}).
			Update("level", position.Level).
			Where(&model.User_Power{
				MemberId: req.ID,
				TeamId:   position.TeamId,
			}).Error

		//删除 用户 和 团队 关联的 旧casbin表
		err = r.DB.Model(&model.Casbin{}).
			Delete(&model.Casbin{}).
			Where(&model.Casbin{
				Ptype: "g",
				V0:    req.ID,
				V1:    position.TeamId,
			}).Error
		if err != nil {
			return err
		}

		//新增 casbin 数据
		if position.Level >= 2 {
			err = r.DB.Model(&model.Casbin{}).
				Create(&model.Casbin{
					Ptype: "g",
					V0:    req.ID,
					V1:    position.TeamId,
					V2:    NormalManger,
				}).Error
			if err != nil {
				return err
			}
		}
		if position.Level == 3 {
			err = r.DB.Model(&model.Casbin{}).
				Create(&model.Casbin{
					Ptype: "g",
					V0:    req.ID,
					V1:    position.TeamId,
					V2:    SuperManger,
				}).Error
			if err != nil {
				return err
			}
		}

	}
	return nil
}
