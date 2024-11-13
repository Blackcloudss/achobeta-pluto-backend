package model

// 个人详细信息表
type Member struct {
	CommonModel
	Name       string `gorm:"column:name; type:varchar(20); index; comment:'真实姓名'"`
	Sex        string `gorm:"column:sex; type:char(2); comment:'性别'"`
	CreateDate string `gorm:"column:create_data; type:char(20); comment:'性别'"`
	IdCard     string `gorm:"column:idcard; type:varchar(50); comment:'身份证'"`
	PhoneNum   uint64 `gorm:"column:phone_num; type:unsigned int; index; comment:'手机号码'"`
	Email      string `gorm:"column:email; type:varchar(30); comment:'邮箱'"`
	Grade      uint   `gorm:"column:grade; type:unsigned int; index; comment:'年级'"`
	Major      string `gorm:"column:major; type:varchar(30); index; comment:'专业'"`
	StudentID  uint64 `gorm:"column:student_id; type:unsigned int; comment:'学号'"`
	Experience string `gorm:"column:experience; type:varchar(200); comment:'实习、创业、就职经历'"`
	Status     string `gorm:"column:status; type:varchar(10); index; comment:'现状'"`
	LikeCount  uint64 `gorm:"column:like_count; type:unsigned int; comment:'点赞数量'"`
}

func (t *Member) TableName() string {
	return "member"
}

// 点赞表
type Like_Status struct {
	CommonModel
	IsLiked         bool  `gorm:"column:is_liked; type:unsigned int; index; comment:'点赞情况'"`
	MemberId_Like   int64 `gorm:"column:memberid_like; type:unsigned int; index; comment:'点赞的用户id'"`
	MemberId_BeLike int64 `gorm:"column:memberid_belike; type:unsigned int; index; comment:'被点赞的用户id'"`
}

func (t *Like_Status) TableName() string {
	return "like_status"
}
