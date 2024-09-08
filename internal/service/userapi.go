package service

import (
	"context"
	"errors"
	pb "kratos-realworld/api/user/v1"
	"kratos-realworld/internal/biz"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/middleware/auth"
	"kratos-realworld/internal/util"
	"time"

	"github.com/jinzhu/copier"
)

type UserApiService struct {
	pb.UnimplementedUserApiServer
	conf *conf.Server
	uc   *biz.UserUsecase
}

func NewUserApiService(uc *biz.UserUsecase, conf *conf.Server) *UserApiService {
	return &UserApiService{uc: uc, conf: conf}
}

func (u *UserApiService) generateToken(userId int64) string {
	return auth.GenerateToken(auth.NewTokenClaims(userId, time.Hour), u.conf.Auth.Jwt.Secret)
}

func (u *UserApiService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReply, error) {
	// 密码加密
	passwordHash, err := util.GenerateFromPassword(req.User.Password)
	if err != nil {
		return nil, err
	}

	user, err := u.uc.CreateUser(ctx, &biz.User{
		Username: req.User.Username,
		Email:    req.User.Email,
		Password: passwordHash,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UserReply{
		User: &pb.User{
			Username: user.Username,
			Email:    user.Email,
			Token:    u.generateToken(user.ID),
		},
	}, nil
}

func (s *UserApiService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReply, error) {

	user, err := s.uc.FindByEmail(ctx, req.User.Email)
	if err != nil {
		return nil, err
	}

	if !util.CompareHashAndPassword(user.Password, req.User.Password) {
		return nil, errors.New("User Not Found")
	}

	var u pb.User
	if err = copier.Copy(&u, user); err != nil {
		return nil, err
	}

	u.Token = s.generateToken(user.ID)
	return &pb.UserReply{
		User: &u,
	}, nil
}
