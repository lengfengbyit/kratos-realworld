package service

import (
	"context"
	"errors"
	pb "kratos-realworld/api/article/v1"
	"kratos-realworld/internal/biz"
	"kratos-realworld/internal/middleware/auth"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type ArticleService struct {
	pb.UnimplementedArticleServer

	biz        *biz.ArticleUsecase
	profileBiz *biz.ProfileUsecase
	log        *log.Helper
}

func NewArticleService(biz *biz.ArticleUsecase, profileBiz *biz.ProfileUsecase, logger log.Logger) *ArticleService {
	return &ArticleService{
		biz:        biz,
		profileBiz: profileBiz,
		log:        log.NewHelper(logger),
	}
}

func (s *ArticleService) ListArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	if req.Limit == 0 {
		req.Limit = 10
	}

	var artParams biz.Article
	if err := copier.Copy(&artParams, req); err != nil {
		return nil, err
	}

	if req.Author != "" {
		profile, err := s.profileBiz.FindByUsername(ctx, req.Author)
		if err != nil {
			return nil, err
		}
		artParams.AuthorId = profile.ID
	}

	articles, err := s.biz.ListArticle(ctx, &artParams, req.Tag, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	return s.appendAuthors(ctx, articles)
}
func (s *ArticleService) GetArticle(ctx context.Context, req *pb.SlugRequest) (*pb.ArticleReply, error) {
	article, err := s.biz.GetArticle(ctx, req.Slug)
	if err != nil {
		return nil, err
	}

	if article == nil {
		return nil, errors.New("article not found")
	}

	var articleReply pb.ArticleReply
	if err = copier.Copy(&articleReply, article); err != nil {
		return nil, err
	}

	// 日期格式化
	articleReply.CreatedAt = article.CreatedAt.Format(time.DateTime)
	articleReply.UpdatedAt = article.UpdatedAt.Format(time.DateTime)

	return s.appendAuthor(ctx, article.AuthorId, &articleReply)
}
func (s *ArticleService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.ArticleReply, error) {
	var articleParams biz.Article
	if err := copier.Copy(&articleParams, req.Article); err != nil {
		return nil, err
	}

	userId, err := auth.GetUserId(ctx)
	if err != nil {
		return nil, errors.Join(err, errors.New("invalid token"))
	}

	articleParams.AuthorId = userId
	article, err := s.biz.CreateArticle(ctx, &articleParams)
	if err != nil {
		return nil, err
	}

	var articleReply pb.ArticleReply
	if err = copier.Copy(&articleReply, article); err != nil {
		return nil, err
	}

	// 日期格式化
	articleReply.CreatedAt = article.CreatedAt.Format(time.DateTime)
	articleReply.UpdatedAt = article.UpdatedAt.Format(time.DateTime)

	// 获取作者信息
	return s.appendAuthor(ctx, article.AuthorId, &articleReply)
}
func (s *ArticleService) UpdateArticle(ctx context.Context, req *pb.SaveArticleRequest) (*pb.ArticleReply, error) {
	var articleParams biz.Article
	if err := copier.Copy(&articleParams, req.Article); err != nil {
		return nil, err
	}
	_, err := s.biz.UpdateArticle(ctx, &articleParams)
	if err != nil {
		return nil, err
	}

	return s.GetArticle(ctx, &pb.SlugRequest{Slug: req.Slug})
}
func (s *ArticleService) DeleteArticle(ctx context.Context, req *pb.SlugRequest) (*pb.EmptyReply, error) {
	_, _ = s.biz.DeleteArticle(ctx, req.Slug)
	return &pb.EmptyReply{}, nil
}
func (s *ArticleService) FeedArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	return &pb.ListArticleReply{}, nil
}

func (s *ArticleService) GetAuthor(ctx context.Context, userId int64) (*pb.Author, error) {
	user, err := s.profileBiz.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &pb.Author{
		Username:  user.Username,
		Bio:       user.Bio,
		Image:     user.Image,
		Following: user.Following,
	}, nil
}

func (s *ArticleService) appendAuthor(ctx context.Context, authorId int64, articleReply *pb.ArticleReply) (*pb.ArticleReply, error) {
	author, err := s.GetAuthor(ctx, authorId)
	if err != nil {
		log.Errorf("failed to get author: %v", err)
		return articleReply, nil
	}
	articleReply.Author = author

	return articleReply, nil
}

func (s *ArticleService) appendAuthors(ctx context.Context, articles []*biz.Article) (*pb.ListArticleReply, error) {
	var err error
	var articlesReply = make([]*pb.ArticleReply, len(articles))
	for i, article := range articles {
		var reply pb.ArticleReply
		if err = copier.Copy(&reply, article); err != nil {
			return nil, err
		}
		author, err := s.GetAuthor(ctx, article.AuthorId)
		if err != nil {
			s.log.Errorf("failed to get author: %v", err)
		}
		reply.Author = author
		articlesReply[i] = &reply
	}

	return &pb.ListArticleReply{Articles: articlesReply}, nil
}
