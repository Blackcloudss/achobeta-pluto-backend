package repo

import (
	"errors"
	"gorm.io/gorm"
	"tgwp/internal/model"
	"tgwp/internal/types"
	"tgwp/log/zlog"
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
var (
	TRUE  int8 = 1
	FALSE int8 = 0
)

func (r *LikeCountRepo) PutLikeCount(UserId, MemberId int64) (*types.LikeCountResp, error) {
	var IsLiked int8 = 0

	tx := r.DB.Begin() //开启事务 -- 保证数据的原子性

	//查询 用户对该成员的点赞情况
	err := tx.Model(&model.Like_Status{}).
		Select("is_liked").
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
			err = tx.Model(&model.Like_Status{}).
				Create(&model.Like_Status{
					MemberId_Like:   UserId,
					MemberId_BeLike: MemberId,
					IsLiked:         FALSE,
				}).Error
			if err != nil {
				tx.Rollback() //回滚事务
				zlog.Errorf("初始化点赞表失败：%v", err)
				return nil, err
			}
		} else {
			// 其他错误情况
			tx.Rollback()
			zlog.Errorf("对该成员点赞情况查询失败：%v", err)
			return nil, err
		}
	}

	var likecount uint64
	err = tx.Model(&model.Member{}).
		Where(&model.Member{
			CommonModel: model.CommonModel{
				ID: MemberId,
			},
		}).
		Select("like_count").
		First(&likecount).
		Error

	if err != nil {
		tx.Rollback()
		zlog.Errorf("对该成员的点赞数量查询失败%v", err)
		return nil, err
	}

	if IsLiked == FALSE {
		//用户还没有给该成员点赞
		err = tx.Model(&model.Like_Status{}).
			Where(&model.Like_Status{
				MemberId_Like:   UserId,
				MemberId_BeLike: MemberId,
			}).
			Update("is_liked", TRUE).
			Error
		if err != nil {
			tx.Rollback()
			zlog.Errorf("当前用户对该成员修改点赞情况失败%v", err)
			return nil, err
		}

		likecount++

		err = tx.Model(&model.Member{}).
			Where(&model.Member{
				CommonModel: model.CommonModel{
					ID: MemberId,
				},
			}).
			Update("like_count", likecount).Error
		if err != nil {
			tx.Rollback()
			zlog.Errorf("当前用户对该成员点赞失败%v", err)
			return nil, err
		}
	}

	if IsLiked == TRUE {
		//用户已经给该成员点赞
		err = tx.Model(&model.Like_Status{}).
			Where(&model.Like_Status{
				MemberId_Like:   UserId,
				MemberId_BeLike: MemberId,
			}).
			Update("is_liked", FALSE).Error
		if err != nil {
			tx.Rollback()
			zlog.Errorf("当前用户对该成员修改点赞情况失败%v", err)
			return nil, err
		}

		likecount--

		err = tx.Model(&model.Member{}).
			Where(&model.Member{
				CommonModel: model.CommonModel{
					ID: MemberId,
				},
			}).
			Update("like_count", likecount).Error
		if err != nil {
			tx.Rollback()
			zlog.Errorf("当前用户对该成员取消赞失败%v", err)
			return nil, err
		}
	}
	err = tx.Commit().Error // 提交事务
	if err != nil {
		tx.Rollback()
		zlog.Errorf("事务提交失败：%v", err)
		return nil, err
	}
	return &types.LikeCountResp{LikeCount: likecount}, nil
}
