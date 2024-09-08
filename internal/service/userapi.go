package service

import (
	"context"
	"errors"
	pb "kratos-realworld/api/user/v1"
	"kratos-realworld/internal/biz"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/middleware/auth"
	"kratos-realworld/internal/util"
	"log"
	"time"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserApiService struct {
	pb.UnimplementedUserApiServer
	conf *conf.Server
	uc   *biz.UserUsecase
}

func NewUserApiService(uc *biz.UserUsecase, conf *conf.Server) *UserApiService {
	return &UserApiService{uc: uc, conf: conf}
}

func (s *UserApiService) generateToken(userId int64) string {
	duration, err := time.ParseDuration(s.conf.Auth.Jwt.Expires)
	if err != nil {
		log.Printf("parse duration failed: %v", err)
		duration = time.Hour
	}
	return auth.GenerateToken(auth.NewTokenClaims(userId, duration), s.conf.Auth.Jwt.Secret)
}

func (s *UserApiService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReply, error) {
	// 密码加密
	passwordHash, err := util.GenerateFromPassword(req.User.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.uc.CreateUser(ctx, &biz.User{
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
			Token:    s.generateToken(user.ID),
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

func (s *UserApiService) CurrentUser(ctx context.Context, e *emptypb.Empty) (*pb.UserReply, error) {

	userId, err := auth.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	user, err := s.uc.FindById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return userModelToUserReply(user)
}

func (s *UserApiService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserReply, error) {
	userId, err := auth.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	var user biz.User
	if err = copier.Copy(&user, req.User); err != nil {
		return nil, err
	}

	user.ID = userId

	u, err := s.uc.UpdateUser(ctx, &user)
	if err != nil {
		return nil, err
	}

	return userModelToUserReply(u)
}

func userModelToUserReply(user *biz.User) (*pb.UserReply, error) {
	var u pb.User
	if err := copier.Copy(&u, user); err != nil {
		return nil, err
	}
	return &pb.UserReply{User: &u}, nil
}
