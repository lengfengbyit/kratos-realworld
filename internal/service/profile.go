package service

import (
	"context"
	pb "kratos-realworld/api/profile/v1"
	"kratos-realworld/internal/biz"

	"github.com/jinzhu/copier"
)

type ProfileService struct {
	pb.UnimplementedProfileServer

	pu *biz.ProfileUsecase
}

func NewProfileService(pu *biz.ProfileUsecase) *ProfileService {
	return &ProfileService{pu: pu}
}

func (s *ProfileService) GetProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileReply, error) {
	info, err := s.pu.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return profileModelToProfileReply(info)
}
func (s *ProfileService) Follow(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileReply, error) {
	follow, err := s.pu.Follow(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return profileModelToProfileReply(follow)
}
func (s *ProfileService) Unfollow(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileReply, error) {
	follow, err := s.pu.UnFollow(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return profileModelToProfileReply(follow)
}

func profileModelToProfileReply(info *biz.Profile) (*pb.ProfileReply, error) {
	var data pb.Data
	if err := copier.Copy(&data, info); err != nil {
		return nil, err
	}
	return &pb.ProfileReply{Profile: &data}, nil
}
