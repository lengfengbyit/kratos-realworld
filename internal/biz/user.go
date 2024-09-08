package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// User is a User model.
type User struct {
	ID        int64
	Email     string
	Username  string
	Password  string
	Bio       string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// UserRepo is a User repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByEmail(context.Context, string) (*User, error)
	ListAll(context.Context) ([]*User, error)
	Login(ctx context.Context, user *User) (*User, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateUser creates a User, and returns the new User.
func (uc *UserUsecase) CreateUser(ctx context.Context, model *User) (*User, error) {
	return uc.repo.Save(ctx, model)
}

func (uc *UserUsecase) FindByEmail(ctx context.Context, email string) (*User, error) {
	return uc.repo.FindByEmail(ctx, email)
}
