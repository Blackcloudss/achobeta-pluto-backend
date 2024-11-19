package repo

import (
	"gorm.io/gorm"
	"tgwp/internal/model"
	"tgwp/util"
	"time"
)

// 留给淳桂 -- 通过手机号判断是游客还是团队成员
type JudgeUserRepo struct {
	DB *gorm.DB
}

func NewJudgeUserRepo(db *gorm.DB) *JudgeUserRepo {
	return &JudgeUserRepo{DB: db}
}

func (r JudgeUserRepo) JudgeUser(Phone uint64) (int64, bool, error) {
	defer util.RecordTime(time.Now())()

	var UserId int64

	err := r.DB.Model(&model.Member{}).
		Select(C_Id).
		Where(&model.Member{
			PhoneNum: Phone,
		}).
		First(&UserId).
		Error
	if err != nil {
		return 0, false, err
	}
	if UserId == 0 {
		return 0, false, nil
	}
	return UserId, true, nil
}
