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

// FindByEmail finds the Profile by username.
func (uc *ProfileUsecase) FindByUsername(ctx context.Context, username string) (*Profile, error) {
	userId, beUser, err := uc.getUsers(ctx, username)
	if err != nil {
		return nil, err
	}

	// 检查当前用户是否关注了被查询的用户
	following := uc.followRepo.IsFollowing(ctx, userId, beUser.ID)
	return &Profile{
		ID:        beUser.ID,
		Image:     beUser.Image,
		Bio:       beUser.Bio,
		Username:  beUser.Username,
		Following: following,
	}, nil
}

func (uc *ProfileUsecase) Follow(ctx context.Context, username string) (*Profile, error) {
	userId, beUser, err := uc.getUsers(ctx, username)
	if err != nil {
		return nil, err
	}

	_ = uc.followRepo.Follow(ctx, userId, beUser.ID)
	following := uc.followRepo.IsFollowing(ctx, userId, beUser.ID)
	return &Profile{
		ID:        beUser.ID,
		Image:     beUser.Image,
		Bio:       beUser.Bio,
		Username:  beUser.Username,
		Following: following,
	}, nil
}

func (uc *ProfileUsecase) UnFollow(ctx context.Context, username string) (*Profile, error) {
	userId, beUser, err := uc.getUsers(ctx, username)
	if err != nil {
		return nil, err
	}

	_ = uc.followRepo.Unfollow(ctx, userId, beUser.ID)
	following := uc.followRepo.IsFollowing(ctx, userId, beUser.ID)
	return &Profile{
		ID:        beUser.ID,
		Image:     beUser.Image,
		Bio:       beUser.Bio,
		Username:  beUser.Username,
		Following: following,
	}, nil

}

func (uc *ProfileUsecase) getUsers(ctx context.Context, username string) (userId int64, beUser *User, err error) {
	// 要查询的用户信息
	beUser, err = uc.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return
	}

	// 当前用户 ID
	userId, err = auth.GetUserId(ctx)
	return
}
