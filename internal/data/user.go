package data

import (
	"context"
	"kratos-realworld/internal/biz"
	"kratos-realworld/internal/data/ent/user"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *userRepo) Login(ctx context.Context, user *biz.User) (*biz.User, error) {
	return user, nil
}

func (r *userRepo) Save(ctx context.Context, g *biz.User) (*biz.User, error) {
	info, err := r.data.db.User.
		Create().
		SetUsername(g.Username).
		SetPassword(g.Password).
		SetEmail(g.Email).
		Save(ctx)

	if err != nil {
		r.log.Errorf("User Save error: %v", err)
		return nil, err
	}

	var u biz.User
	if err = copier.Copy(&u, info); err != nil {
		r.log.Errorf("User Save copy error: %v", err)
		return nil, err
	}
	return &u, err
}

func (r *userRepo) Update(ctx context.Context, g *biz.User) (*biz.User, error) {
	return g, nil
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*biz.User, error) {
	info, err := r.data.db.User.Query().Where(user.EmailEQ(email)).First(ctx)
	if err != nil {
		return nil, err
	}

	var u biz.User
	if err = copier.Copy(&u, info); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) ListByHello(context.Context, string) ([]*biz.User, error) {
	return nil, nil
}

func (r *userRepo) ListAll(context.Context) ([]*biz.User, error) {
	return nil, nil
}
