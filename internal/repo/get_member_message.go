package repo

import (
	"gorm.io/gorm"
	"strings"
	"tgwp/internal/model"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

type MemberDetailRepo struct {
	DB *gorm.DB
}

func NewMemberDetailRepo(db *gorm.DB) *MemberDetailRepo {
	return &MemberDetailRepo{
		DB: db,
	}
}

// GetMemberDetail
//
//	@Description: 查询 团队成员详细信息
//	@receiver r
//	@param userID
//	@return *types.GetMemberDetailResp
//	@return error
func (r *MemberDetailRepo) GetMemberDetail(userID int64) (*types.GetMemberDetailResp, error) {
	var resp types.GetMemberDetailResp

	var positions []struct {
		TeamID   int64  `json:"team_id"`
		TeamName string `json:"team_name"`
		NodeID   int64  `json:"node_id"`
		NodeName string `json:"node_name"`
		Level    int    `json:"level"`
	}

	// 查询用户的基本信息
	err := r.DB.Model(&model.Member{}).
		Select(`member.name, member.sex, member.create_date, member.id_card, 
				member.phone_num, member.email, member.grade, member.major, 
				member.student_id, member.experience, member.status, member.like_count`).
		Where("id = ?", userID).
		First(&resp).
		Error
	if err != nil {
		zlog.Errorf("查询用户信息失败: %v", err)
		return nil, err
	}

	// 查询用户的团队、职位和权限
	err = r.DB.Model(&model.Team_Member_Structure{}).
		Select(`team.id AS team_id, team.name AS team_name, 
				structure.id AS node_id, structure.node_name AS node_name, 
				user_power.level AS level`).
		Joins("JOIN team ON team.id = team_member_structure.team_id").
		Joins("JOIN structure ON structure.id = team_member_structure.structure_id").
		Joins("JOIN user_power ON user_power.member_id = team_member_structure.member_id AND user_power.team_id = team_member_structure.team_id").
		Where("team_member_structure.member_id = ?", userID).
		Find(&positions).
		Error
	if err != nil {
		zlog.Errorf("查询团队职位信息失败: %w", err)
		return nil, err
	}

	// 整理结果：按团队分组职位
	positionMap := make(map[int64]*types.MemberPositions)
	for _, pos := range positions {
		//将职位数据按团队分组
		if _, exists := positionMap[pos.TeamID]; !exists {
			positionMap[pos.TeamID] = &types.MemberPositions{
				TeamId:      pos.TeamID,
				TeamName:    pos.TeamName,
				Level:       pos.Level,
				StructNodes: []types.StructureNodes{},
			}
		}
		positionMap[pos.TeamID].StructNodes = append(positionMap[pos.TeamID].StructNodes, types.StructureNodes{
			ID:   pos.NodeID,
			Name: pos.NodeName,
		})
	}

	// 转换为切片
	for _, pos := range positionMap {
		resp.MemberPosition = append(resp.MemberPosition, *pos)
	}

	return &resp, nil
}

type MemberlistRepo struct {
	DB *gorm.DB
}

func NewMemberlistRepo(db *gorm.DB) *MemberlistRepo {
	return &MemberlistRepo{DB: db}
}

// MemberlistRepo
//
//	@Description: 分页查询  用户基本信息
//	@receiver r
//	@param TeamId
//	@param Page
//	@param Perpage
//	@return types.MemberlistResp
//	@return error
func (r *MemberlistRepo) MemberlistRepo(TeamId int64, Page, Perpage int) (types.MemberlistResp, error) {

	var Members []types.Members

	Offset := (Page - 1) * Perpage
	if Offset < 0 {
		Offset = 0
	}

	//先这样，要美观的话之后更改
	err := r.DB.Table("member").
		Select(`member.id AS member_id, member.name, member.grade, member.major, member.status, 
			member.phone_num, GROUP_CONCAT(structure.node_name) AS positions`).
		Joins(`JOIN team_member_structure ON team_member_structure.member_id = member.id`).
		Joins(`JOIN structure ON structure.id = team_member_structure.structure_id`).
		Where("team_member_structure.team_id = ?", TeamId).
		Group("member.id"). // 按 Member 分组
		Offset(Offset).
		Limit(Perpage).
		Find(&Members).
		Error

	//在查询后将 Positions 的逗号分隔值拆分为切片： 查询 当前用户所处团队的每个职位名称
	for i, member := range Members {
		Members[i].PositionList = strings.Split(member.Positions, ",")
	}

	if err != nil {
		return types.MemberlistResp{}, err
	}
	return types.MemberlistResp{Members: Members}, nil
}
