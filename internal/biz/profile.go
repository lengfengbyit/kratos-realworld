package biz

import (
	"context"
	"kratos-realworld/internal/middleware/auth"

	"github.com/go-kratos/kratos/v2/log"
)

// Profile is a Profile model.
type Profile struct {
	ID        int64
	Username  string
	Bio       string
	Image     string
	Following bool
}

// ProfileRepo is a Profile repo.
type ProfileRepo interface {
}

// ProfileUsecase is a Profile usecase.
type ProfileUsecase struct {
	userRepo   UserRepo
	followRepo FollowRepo
	log        *log.Helper
}

// NewProfileUsecase new a Profile usecase.
func NewProfileUsecase(repo UserRepo, followRepo FollowRepo, logger log.Logger) *ProfileUsecase {
	return &ProfileUsecase{userRepo: repo, followRepo: followRepo, log: log.NewHelper(logger)}
}

func (biz *ProfileUsecase) FindByUserId(ctx context.Context, beUserId int64) (*Profile, error) {

	beUser, err := biz.userRepo.FindById(ctx, beUserId)
	if err != nil {
		return nil, err
	}

	// 当前用户 ID
	userId, err := auth.GetUserId(ctx)

	// 检查当前用户是否关注了被查询的用户
	following := biz.followRepo.IsFollowing(ctx, userId, beUser.ID)
	return &Profile{
		ID:        beUser.ID,
		Image:     beUser.Image,
		Bio:       beUser.Bio,
		Username:  beUser.Username,
		Following: following,
	}, nil
}

// FindByEmail finds the Profile by username.
func (biz *ProfileUsecase) FindByUsername(ctx context.Context, username string) (*Profile, error) {
	userId, beUser, err := biz.getUsers(ctx, username)
	if err != nil {
		return nil, err
	}

	// 检查当前用户是否关注了被查询的用户
	following := biz.followRepo.IsFollowing(ctx, userId, beUser.ID)
	return &Profile{
		ID:        beUser.ID,
		Image:     beUser.Image,
		Bio:       beUser.Bio,
		Username:  beUser.Username,
		Following: following,
	}, nil
}

func (biz *ProfileUsecase) Follow(ctx context.Context, username string) (*Profile, error) {
	userId, beUser, err := biz.getUsers(ctx, username)
	if err != nil {
		return nil, err
	}

	_ = biz.followRepo.Follow(ctx, userId, beUser.ID)
	following := biz.followRepo.IsFollowing(ctx, userId, beUser.ID)
	return &Profile{
		ID:        beUser.ID,
		Image:     beUser.Image,
		Bio:       beUser.Bio,
		Username:  beUser.Username,
		Following: following,
	}, nil
}

func (biz *ProfileUsecase) UnFollow(ctx context.Context, username string) (*Profile, error) {
	userId, beUser, err := biz.getUsers(ctx, username)
	if err != nil {
		return nil, err
	}

	_ = biz.followRepo.Unfollow(ctx, userId, beUser.ID)
	following := biz.followRepo.IsFollowing(ctx, userId, beUser.ID)
	return &Profile{
		ID:        beUser.ID,
		Image:     beUser.Image,
		Bio:       beUser.Bio,
		Username:  beUser.Username,
		Following: following,
	}, nil

}

func (biz *ProfileUsecase) getUsers(ctx context.Context, username string) (userId int64, beUser *User, err error) {
	// 要查询的用户信息
	beUser, err = biz.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return
	}

	// 当前用户 ID
	userId, err = auth.GetUserId(ctx)
	return
}
