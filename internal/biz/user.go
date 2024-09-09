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
	Update(context.Context, *User) error
	FindByEmail(context.Context, string) (*User, error)
	FindByUsername(context.Context, string) (*User, error)
	FindIdByUsername(context.Context, string) (int64, error)
	FindById(context.Context, int64) (*User, error)
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

// FindByEmail finds the User by Email.
func (uc *UserUsecase) FindByEmail(ctx context.Context, email string) (*User, error) {
	return uc.repo.FindByEmail(ctx, email)
}

// FindByUsername finds the User by Email.
func (uc *UserUsecase) FindByUsername(ctx context.Context, email string) (*User, error) {
	return uc.repo.FindByUsername(ctx, email)
}

// FindById finds the User by ID.
func (uc *UserUsecase) FindById(ctx context.Context, id int64) (*User, error) {
	return uc.repo.FindById(ctx, id)
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, model *User) (*User, error) {
	err := uc.repo.Update(ctx, model)
	if err != nil {
		return nil, err
	}
	return uc.repo.FindById(ctx, model.ID)
}
