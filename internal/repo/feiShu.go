package repo

import (
	"fmt"
	"gorm.io/gorm"
	"tgwp/internal/handler"
	"tgwp/internal/model"
	"tgwp/internal/response"
	"tgwp/log/zlog"
)

type FeiShuRepo struct {
	db *gorm.DB
}

func NewFeiShuRepo(db *gorm.DB) *FeiShuRepo {
	return &FeiShuRepo{db: db}
}

// GetFeiShuOpenID
//
//	@Description: 获取FeiShuOpenID
//	@receiver r
//	@param UserID
//	@return OpenID
//	@return err
func (r FeiShuRepo) GetFeiShuOpenID(UserID int64) (OpenID string, err error) {
	member := model.Member{}

	err = r.db.Model(&model.Member{}).Where("id =?", UserID).Order("created_at desc").First(&member).Error
	if err != nil {
		zlog.Errorf("get member err:%v", err)
		err = response.ErrResp(err, response.DATABASE_ERROR)
		return
	}
	// 如果用户表里还没有FeiShuOpenID，则获取并保存
	if len(member.FeiShuOpenID) <= 0 {
		OpenID, err = handler.GetFeiShuUserOpenID(fmt.Sprintf("%d", member.PhoneNum))
		member.FeiShuOpenID = OpenID
		if err != nil {
			zlog.Errorf("get feishu openid err:%v", err)
			err = response.ErrResp(err, response.FEISHU_ERROR)
			return
		}
		err = r.db.Save(&member).Error
		if err != nil {
			zlog.Errorf("save member err:%v", err)
			err = response.ErrResp(err, response.DATABASE_ERROR)
			return
		}
	} else {
		OpenID = member.FeiShuOpenID
	}

	return
}