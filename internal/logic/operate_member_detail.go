package logic

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"tgwp/global"
	"tgwp/internal/repo"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
	"tgwp/util"
	"time"
)

type MemberDetailLogic struct{}

func NewMemberDetailLogic() *MemberDetailLogic {
	return &MemberDetailLogic{}
}

// GetMemberDetail
//
//	@Description: 查看用户详细信息
//	@receiver l
//	@param ctx
//	@param req
//	@return resp
//	@return err
func (l *MemberDetailLogic) GetMemberDetail(ctx context.Context, req types.GetMemberDetailReq) (resp *types.GetMemberDetailResp, err error) {
	defer util.RecordTime(time.Now())()

	resp, err = repo.NewMemberDetailRepo(global.DB).GetMemberDetail(req.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zlog.CtxWarnf(ctx, "user not found: %v", err)
			return nil, response.ErrResp(err, codeUserNotFound)
		} else {
			zlog.Errorf("get user error: %v", err)
			//注意 codeUserFoundField 只是事例，具体根据实际情况定义 response.ErrResp这里用作包装错误和响应，使得错误更加通用
			return nil, response.ErrResp(err, codeUserFoundField)
		}

	}
	return
}

var (
	codeServerError     = response.MsgCode{Code: 60003, Msg: "上锁失败"}
	codeOperationLocked = response.MsgCode{Code: 20013, Msg: "该操作被加了锁"}
)

type LikeCountLogic struct{}

func NewLikeCountLogic() *LikeCountLogic {
	return &LikeCountLogic{}
}

// PutLikeCount
//
//	@Description: 用户点赞/取消赞
//	@receiver l
//	@param ctx
//	@param UserId
//	@param MemberId
//	@return resp
//	@return err
func (l *LikeCountLogic) PutLikeCount(ctx context.Context, UserId, MemberId int64) (resp *types.LikeCountResp, err error) {
	defer util.RecordTime(time.Now())()

	// 用 redis 加锁
	lockKey := fmt.Sprintf("like:lock:user:%d:member:%d", UserId, MemberId)
	// 尝试获取锁   返回布尔值 true:成功获取锁  false:上锁失败
	locked, err := global.Rdb.SetNX(ctx, lockKey, 1, 5*time.Second).Result()
	if err != nil {
		zlog.CtxErrorf(ctx, "Redis 上锁失败: %v", err)
		return nil, response.ErrResp(err, codeServerError)
	}
	if !locked {
		// 未获取到锁，说明该操作正在被其他请求处理
		zlog.CtxInfof(ctx, "Operation is locked for user: %d, member: %d", UserId, MemberId)
		return nil, response.ErrResp(fmt.Errorf("该操作被加了锁"), codeOperationLocked)
	}
	//释放锁
	defer global.Rdb.Del(ctx, lockKey)

	resp, err = repo.NewLikeCountRepo(global.DB).PutLikeCount(UserId, MemberId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zlog.CtxWarnf(ctx, "user not found: %v", err)
			return nil, response.ErrResp(err, codeUserNotFound)
		} else {
			zlog.Errorf("get user error: %v", err)
			//注意 codeUserFoundField 只是事例，具体根据实际情况定义 response.ErrResp这里用作包装错误和响应，使得错误更加通用
			return nil, response.ErrResp(err, codeUserFoundField)
		}
	}
	return
}
