package repo

import (
	"errors"
	"gorm.io/gorm"
	"tgwp/internal/model"
	"tgwp/internal/types"
	"tgwp/log/zlog"
	"time"
)

type LikeCountRepo struct {
	DB *gorm.DB
}

func NewLikeCountRepo(db *gorm.DB) *LikeCountRepo {
	return &LikeCountRepo{
		DB: db,
	}
}

// PutLikeCount
//
//	@Description: 用户点赞/取消赞
//	@receiver r
//	@param UserId
//	@param MemberId
//	@return *types.LikeCountResp
//	@return error
func (r *LikeCountRepo) PutLikeCount(UserId, MemberId int64) (*types.LikeCountResp, error) {
	var IsLiked bool
	//查询 用户对该成员的点赞情况
	err := r.DB.Model(&model.Like_Status{}).
		Where(&model.Like_Status{
			MemberId_Like:   UserId,
			MemberId_BeLike: MemberId,
		}).
		First(&IsLiked).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有找到记录的情况,说明该表还没有相关数据
			// 创建数据
			err = r.DB.Model(&model.Like_Status{}).
				Create(&model.Like_Status{
					CommonModel: model.CommonModel{
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
					MemberId_Like:   UserId,
					MemberId_BeLike: MemberId,
					IsLiked:         false,
				}).Error
			if err != nil {
				zlog.Errorf("初始化点赞表失败：%v", err)
				return nil, err
			}
		} else {
			// 其他错误情况
			zlog.Errorf("对该成员点赞情况查询失败：%v", err)
			return nil, err
		}
	}

	var likecount uint64
	err = r.DB.Model(&model.Member{}).
		Where(&model.Member{
			CommonModel: model.CommonModel{
				ID: MemberId,
			},
		}).
		First(&likecount).
		Error

	if err != nil {
		zlog.Errorf("对该成员的点赞数量查询失败%v", err)
		return nil, err
	}

	if IsLiked == false {
		//用户还没有给该成员点赞
		err = r.DB.Model(&model.Like_Status{}).
			Updates(&model.Like_Status{
				CommonModel: model.CommonModel{
					UpdatedAt: time.Now(),
				},
				IsLiked: true,
			}).
			Where(&model.Like_Status{
				MemberId_Like:   UserId,
				MemberId_BeLike: MemberId,
			}).Error
		if err != nil {
			zlog.Errorf("当前用户对该成员修改点赞情况失败%v", err)
			return nil, err
		}

		likecount++

		err = r.DB.Model(&model.Member{}).
			Updates(&model.Member{
				CommonModel: model.CommonModel{
					UpdatedAt: time.Now(),
				},
				LikeCount: likecount,
			}).
			Where(&model.Member{
				CommonModel: model.CommonModel{
					ID: MemberId,
				},
			}).Error
		if err != nil {
			zlog.Errorf("当前用户对该成员点赞失败%v", err)
			return nil, err
		}
	} else {
		//用户已经给该成员点赞
		err = r.DB.Model(&model.Like_Status{}).
			Updates(&model.Like_Status{
				CommonModel: model.CommonModel{
					UpdatedAt: time.Now(),
				},
				IsLiked: false,
			}).
			Where(&model.Like_Status{
				MemberId_Like:   UserId,
				MemberId_BeLike: MemberId,
			}).Error
		if err != nil {
			zlog.Errorf("当前用户对该成员修改点赞情况失败%v", err)
			return nil, err
		}

		likecount--

		err = r.DB.Model(&model.Member{}).
			Updates(&model.Member{
				CommonModel: model.CommonModel{
					UpdatedAt: time.Now(),
				},
				LikeCount: likecount,
			},
			).
			Where(&model.Member{
				CommonModel: model.CommonModel{
					ID: MemberId,
				},
			}).Error
		if err != nil {
			zlog.Errorf("当前用户对该成员取消失败%v", err)
			return nil, err
		}
	}
	return &types.LikeCountResp{LikeCount: likecount}, nil
}
