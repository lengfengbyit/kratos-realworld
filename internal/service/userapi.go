package service

import (
	"context"
	pb "kratos-realworld/api/user/v1"
	"kratos-realworld/internal/biz"
)

type UserApiService struct {
	pb.UnimplementedUserApiServer

	uc *biz.UserUsecase
}

func NewUserApiService(uc *biz.UserUsecase) *UserApiService {
	return &UserApiService{uc: uc}
}

func (u *UserApiService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReply, error) {
	user, err := u.uc.CreateUser(ctx, &biz.User{
		Username: req.User.Username,
		Email:    req.User.Email,
		Password: req.User.Password,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UserReply{
		User: &pb.User{
			Username: user.Username,
			Email:    user.Email,
			Token:    "token",
		},
	}, nil
}

func (s *UserApiService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReply, error) {
	return &pb.UserReply{
		User: &pb.User{
			Username: "tom",
			Email:    "tom@gmail.com",
			Token:    "token",
		},
	}, nil
}
