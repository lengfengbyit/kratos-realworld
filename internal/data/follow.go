package data

import (
	"context"
	"kratos-realworld/internal/biz"
	"kratos-realworld/internal/data/ent"
	"kratos-realworld/internal/data/ent/follow"
	"kratos-realworld/internal/util"

	"github.com/go-kratos/kratos/v2/log"
)

type FollowRepo struct {
	data *Data
	log  *log.Helper
}

// NewFollowRepo .
func NewFollowRepo(data *Data, logger log.Logger) biz.FollowRepo {
	return &FollowRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (f FollowRepo) Follow(ctx context.Context, userId int64, beUserId int64) bool {
	info, err := f.data.db.Follow.Create().SetUserID(userId).SetBeUserID(beUserId).Save(ctx)
	if err != nil {
		f.log.WithContext(ctx).Errorf("Follow err: %v", err)
		return false
	}
	return info != nil
}

func (f FollowRepo) Unfollow(ctx context.Context, userId int64, beUserId int64) bool {
	_, err := f.data.db.Follow.Delete().Where(follow.And(follow.UserID(userId), follow.BeUserID(beUserId))).Exec(ctx)
	return err == nil
}

func (f FollowRepo) IsFollowing(ctx context.Context, userId int64, beUserId int64) bool {
	exist, err := f.data.db.Follow.Query().Where(follow.And(follow.UserID(userId), follow.BeUserID(beUserId))).Exist(ctx)
	if err != nil {
		f.log.WithContext(ctx).Errorf("IsFollowing err: %v", err)
		return false
	}
	return exist
}

// GetFollowUserIds 获取用户的关注用户id
func (f FollowRepo) GetFollowUserIds(ctx context.Context, userId int64) ([]int64, error) {
	lst, err := f.data.db.Follow.Query().Where(follow.UserID(userId)).Select(follow.FieldBeUserID).All(ctx)
	if err != nil {
		return nil, err
	}

	beUserIds := util.SliceMap(lst, func(item *ent.Follow, index int) int64 {
		return item.BeUserID
	})
	return beUserIds, nil
}
