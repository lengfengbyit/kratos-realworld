package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Favorite struct {
	ArticleId int64
	UserId    int64
}

type FavoriteRepo interface {
	Favorite(ctx context.Context, userId int64, articleId int64) bool
	UnFavorite(ctx context.Context, userId int64, articleId int64) bool
	GetUserIds(ctx context.Context, articleId int64) ([]int64, error)
	GetArticleIds(ctx context.Context, userId int64) ([]int64, error)
	IsFavorite(ctx context.Context, userId int64, articleId int64) bool
	FavoriteCount(ctx context.Context, articleId int64) int
}

type FavoriteUsecase struct {
	repo FavoriteRepo
	log  *log.Helper
}

func NewFavoriteUsecase(repo FavoriteRepo, logger log.Logger) *FavoriteUsecase {
	return &FavoriteUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (f *FavoriteUsecase) Favorite(ctx context.Context, userId int64, articleId int64) bool {
	if f.repo.IsFavorite(ctx, userId, articleId) {
		return true
	}
	return f.repo.Favorite(ctx, userId, articleId)
}
func (f *FavoriteUsecase) UnFavorite(ctx context.Context, userId int64, articleId int64) bool {
	return f.repo.UnFavorite(ctx, userId, articleId)
}

func (f *FavoriteUsecase) FavoriteCount(ctx context.Context, articleId int64) uint32 {
	return uint32(f.repo.FavoriteCount(ctx, articleId))
}
