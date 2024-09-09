package biz

import (
	"context"
)

type Follow struct {
	ID        int64
	Username  string
	Bio       string
	Image     string
	Following bool
}

// FollowRepo is a Follow repo.
type FollowRepo interface {
	IsFollowing(ctx context.Context, userId int64, beUserId int64) bool
	Follow(ctx context.Context, userId int64, beUserId int64) bool
	Unfollow(ctx context.Context, userId int64, beUserId int64) bool
	GetFollowUserIds(ctx context.Context, userId int64) ([]int64, error)
}
