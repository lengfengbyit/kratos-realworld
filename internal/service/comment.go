package service

import (
	"context"
	pb "kratos-realworld/api/comment/v1"
	"kratos-realworld/internal/biz"
	"kratos-realworld/internal/middleware/auth"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type CommentService struct {
	pb.UnimplementedCommentServer

	biz        *biz.CommonUsecase
	articleBiz *biz.ArticleUsecase
	log        *log.Helper
}

func NewCommentService(biz *biz.CommonUsecase, articleBiz *biz.ArticleUsecase, logger log.Logger) *CommentService {
	return &CommentService{
		biz:        biz,
		articleBiz: articleBiz,
		log:        log.NewHelper(logger),
	}
}

func (s *CommentService) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CommentReply, error) {
	// 获取文章 ID
	article, err := s.articleBiz.GetArticle(ctx, req.Slug)
	if err != nil {
		return nil, err
	}

	// 获取当前用户 ID
	userId, err := auth.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	var comm = &biz.Comment{
		ArticleId: article.Id,
		UserId:    userId,
		Body:      req.Comment.Body,
	}

	_, err = s.biz.CreateComment(ctx, comm)
	return &pb.CommentReply{}, err
}
func (s *CommentService) ListComment(ctx context.Context, req *pb.ListCommentRequest) (*pb.ListCommentReply, error) {
	// 获取文章 ID
	article, err := s.articleBiz.GetArticle(ctx, req.Slug)
	if err != nil {
		return nil, err
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	lst, err := s.biz.ListComment(ctx, article.Id, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	var lstReply = make([]*pb.CommentInfoReply, len(lst))
	for i, item := range lst {
		var tmp pb.CommentInfoReply
		if err = copier.Copy(&tmp, item); err != nil {
			return nil, err
		}
		tmp.CreatedAt = item.CreatedAt.Format(time.DateTime)
		lstReply[i] = &tmp
	}

	return &pb.ListCommentReply{Comments: lstReply}, nil
}
func (s *CommentService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentReply, error) {
	err := s.biz.DeleteComment(ctx, req.Id)
	return &pb.DeleteCommentReply{}, err
}
