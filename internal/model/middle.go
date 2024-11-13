package model

// 团队-架构-用户中间表
type Team_Member_Structure struct {
	CommonModel
	TeamId      int64 `gorm:"column:team_id; type:bigint; index; not null; comment:'团队ID'"`
	MemberId    int64 `gorm:"column:member_id; type:bigint; index; not null; comment:'成员ID'"`
	StructureId int64 `gorm:"column:structure_id; type:bigint; index; not null; comment:'职位ID'"`

	// 外键关联
	Team      Team      `gorm:"foreignKey:TeamId; references:ID"`
	Member    Member    `gorm:"foreignKey:MemberId; references:ID"`
	Structure Structure `gorm:"foreignKey:StructureId; references:ID"`
}

func (t *Team_Member_Structure) TableName() string {
	return "team_member_structure"
}

// 用户权限表
type User_Power struct {
	CommonModel
	MemberId int64 `gorm:"column:member_id; type:bigint;   index;not null; comment:'成员ID'"`
	TeamId   int64 `gorm:"column:team_id;   type:bigint;   index;not null; comment:'团队ID'"`
	Level    int   `gorm:"column:level;     type:int; index;not null; default:1; comment:'权限等级'"`

	// 外键关联
	Member Member `gorm:"foreignKey:MemberId; references:ID"`
	Team   Team   `gorm:"foreignKey:TeamId; references:ID"`
}

func (t *User_Power) TableName() string {
	return "user_power"
}
