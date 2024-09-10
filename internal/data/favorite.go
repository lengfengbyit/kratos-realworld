package data

import (
	"context"
	"kratos-realworld/internal/biz"
	"kratos-realworld/internal/data/ent"
	"kratos-realworld/internal/data/ent/favorite"
	"kratos-realworld/internal/util"

	"github.com/go-kratos/kratos/v2/log"
)

type favoriteRepo struct {
	data *Data
	log  *log.Helper
}

func NewFavoriteRepo(data *Data, logger log.Logger) biz.FavoriteRepo {
	return &favoriteRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (f *favoriteRepo) Favorite(ctx context.Context, userId int64, articleId int64) bool {
	_, err := f.data.db.Favorite.Create().SetUserID(userId).SetArticleID(articleId).Save(ctx)
	if err != nil {
		f.log.Errorf("favorite err:%v", err)
		return false
	}
	return true
}

func (f *favoriteRepo) UnFavorite(ctx context.Context, userId int64, articleId int64) bool {
	_, err := f.data.db.Favorite.Delete().
		Where(favorite.And(favorite.UserID(userId), favorite.ArticleID(articleId))).
		Exec(ctx)
	if err != nil {
		f.log.Errorf("unFavorite err:%v", err)
		return false
	}
	return true
}

func (f *favoriteRepo) GetUserIds(ctx context.Context, articleId int64) ([]int64, error) {
	lst, err := f.data.db.Favorite.Query().
		Where(favorite.ArticleID(articleId)).
		Select(favorite.FieldUserID).
		All(ctx)
	if err != nil {
		return nil, err
	}

	userIds := util.SliceMap(lst, func(item *ent.Favorite, index int) int64 {
		return item.UserID
	})
	return userIds, nil
}

func (f *favoriteRepo) GetArticleIds(ctx context.Context, userId int64) ([]int64, error) {
	lst, err := f.data.db.Favorite.Query().
		Where(favorite.UserID(userId)).
		Select(favorite.FieldArticleID).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return util.SliceMap(lst, func(item *ent.Favorite, index int) int64 {
		return item.ArticleID
	}), nil
}

func (f *favoriteRepo) IsFavorite(ctx context.Context, userId int64, articleId int64) bool {
	exist, err := f.data.db.Favorite.Query().Where(favorite.And(favorite.UserID(userId), favorite.ArticleID(articleId))).Exist(ctx)
	if err != nil {
		f.log.Errorf("IsFavorite err:%v", err)
		return false
	}
	return exist
}

func (f *favoriteRepo) FavoriteCount(ctx context.Context, articleId int64) int {
	count, err := f.data.db.Favorite.Query().Where(favorite.ArticleID(articleId)).Count(ctx)
	if err != nil {
		f.log.Errorf("FavoriteCount err:%v", err)
		return 0
	}
	return count
}
