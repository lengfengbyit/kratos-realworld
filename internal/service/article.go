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

	biz         *biz.ArticleUsecase
	profileBiz  *biz.ProfileUsecase
	favoriteBiz *biz.FavoriteUsecase
	log         *log.Helper
}

func NewArticleService(biz *biz.ArticleUsecase, profileBiz *biz.ProfileUsecase, favoriteBiz *biz.FavoriteUsecase, logger log.Logger) *ArticleService {
	return &ArticleService{
		biz:         biz,
		profileBiz:  profileBiz,
		favoriteBiz: favoriteBiz,
		log:         log.NewHelper(logger),
	}
}

func (s *ArticleService) ListArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	return s.listArticle(ctx, req, nil)
}

func (s *ArticleService) listArticle(ctx context.Context, req *pb.ListArticleRequest, beUserIds []int64) (*pb.ListArticleReply, error) {
	if req.Limit == 0 {
		req.Limit = 10
	}

	var artParams biz.ListArticleRequest
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

	if req.Favorited != "" {
		profile, err := s.profileBiz.FindByUsername(ctx, req.Favorited)
		if err != nil {
			return nil, err
		}
		artParams.UserIds = []int64{profile.ID}
	}

	if beUserIds != nil {
		artParams.UserIds = beUserIds
	}

	total, articles, err := s.biz.ListArticle(ctx, &artParams)
	if err != nil {
		return nil, err
	}

	var articlesReply = make([]*pb.ArticleReply, len(articles))
	if err = copier.Copy(&articlesReply, articles); err != nil {
		return nil, err
	}

	return &pb.ListArticleReply{Articles: articlesReply, ArticlesCount: uint32(total)}, nil
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

	// 检查当前用户是否点赞
	userId, err := auth.GetUserId(ctx)
	if err == nil {
		if s.favoriteBiz.Favorite(ctx, userId, article.Id) {
			articleReply.Favorited = true
		}
	}

	// 获取文章的点赞数
	articleReply.FavoritesCount = s.favoriteBiz.FavoriteCount(ctx, article.Id)

	return &articleReply, nil
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
	article, err := s.GetArticle(ctx, &pb.SlugRequest{Slug: req.Slug})
	if err != nil {
		return nil, err
	}

	var articleParams biz.Article
	if err = copier.Copy(&articleParams, req.Article); err != nil {
		return nil, err
	}
	_, err = s.biz.UpdateArticle(ctx, &articleParams)
	if err != nil {
		return nil, err
	}

	if err = copier.Copy(article, req.Article); err != nil {
		return nil, err
	}

	return s.GetArticle(ctx, &pb.SlugRequest{Slug: req.Slug})
}
func (s *ArticleService) DeleteArticle(ctx context.Context, req *pb.SlugRequest) (*pb.EmptyReply, error) {
	_, _ = s.biz.DeleteArticle(ctx, req.Slug)
	return &pb.EmptyReply{}, nil
}

// FeedArticle 获取关注用户的文章列表
func (s *ArticleService) FeedArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	// 获取当前用户 id
	userId, err := auth.GetUserId(ctx)
	if err != nil {
		return nil, errors.Join(err, errors.New("invalid token"))
	}

	// 获取关注用户Ids
	followUserIds, err := s.profileBiz.GetFollowUserIds(ctx, userId)

	if err != nil {
		return nil, err
	}

	// 获取文章列表
	return s.listArticle(ctx, req, followUserIds)
}

// FavoriteArticle 点赞文章
func (s *ArticleService) FavoriteArticle(ctx context.Context, req *pb.SlugRequest) (*pb.ArticleReply, error) {
	userId, articleId, err := s.getUserIdAndArticleId(ctx, req)
	if err != nil {
		return nil, err
	}

	// 保存点赞关系
	if s.favoriteBiz.Favorite(ctx, userId, articleId) {
		return s.GetArticle(ctx, req)
	}

	return nil, errors.New("failed to favorite article")
}

// UnFavoriteArticle 取消点赞文章
func (s *ArticleService) UnFavoriteArticle(ctx context.Context, req *pb.SlugRequest) (*pb.ArticleReply, error) {
	userId, articleId, err := s.getUserIdAndArticleId(ctx, req)
	if err != nil {
		return nil, err
	}

	// 保存点赞关系
	if s.favoriteBiz.UnFavorite(ctx, userId, articleId) {
		return s.GetArticle(ctx, req)
	}

	return nil, errors.New("failed to unFavorite article")
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

func (s *ArticleService) getUserIdAndArticleId(ctx context.Context, req *pb.SlugRequest) (int64, int64, error) {
	// 获取用户 ID
	userId, err := auth.GetUserId(ctx)
	if err != nil {
		return 0, 0, errors.Join(err, errors.New("invalid token"))
	}

	// 获取文章 ID
	article, err := s.biz.GetArticle(ctx, req.Slug)
	if err != nil {
		return 0, 0, err
	}
	return userId, article.Id, nil
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
