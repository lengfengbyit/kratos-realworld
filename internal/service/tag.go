package service

import (
	"context"
	pb "kratos-realworld/api/tag/v1"
	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type TagService struct {
	pb.UnimplementedTagServer

	biz *biz.TagUsecase
	log *log.Helper
}

func NewTagService(biz *biz.TagUsecase, logger log.Logger) *TagService {
	return &TagService{biz: biz, log: log.NewHelper(logger)}
}

func (s *TagService) ListTag(ctx context.Context, req *pb.ListTagRequest) (*pb.ListTagReply, error) {
	tags, err := s.biz.ListTag(ctx, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}
	return &pb.ListTagReply{Tags: tags}, nil
}
