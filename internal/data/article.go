package data

import (
	"context"
	"kratos-realworld/internal/biz"
	"kratos-realworld/internal/data/ent"
	"kratos-realworld/internal/data/ent/article"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *articleRepo) CreateArticle(ctx context.Context, g *biz.Article) (*biz.Article, error) {
	info, err := repo.data.db.Article.Create().
		SetBody(g.Body).
		SetDescription(g.Description).
		SetTitle(g.Title).
		SetAuthorID(g.AuthorId).
		SetTags(strings.Join(g.TagList, ",")).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	var art biz.Article
	if err = copier.Copy(&art, info); err != nil {
		return nil, err
	}

	return &art, err
}

func (repo *articleRepo) UpdateArticle(ctx context.Context, g *biz.Article) (int, error) {
	orm := repo.data.db.Article.Update()

	if g.Body != "" {
		orm.SetBody(g.Body)
	}
	if g.Description != "" {
		orm.SetDescription(g.Description)
	}
	if g.Title != "" {
		orm.SetTitle(g.Title)
	}

	return orm.Save(ctx)
}

func (repo *articleRepo) GetArticle(ctx context.Context, slug uuid.UUID) (*biz.Article, error) {
	info, err := repo.data.db.Article.Query().WithAuthor().Where(article.SlugEQ(slug)).First(ctx)
	if err != nil {
		repo.log.Errorf("get article error: %v", err)
		return nil, err
	}

	var art biz.Article
	if err = copier.Copy(&art, info); err != nil {
		return nil, err
	}
	return &art, err
}

func (repo *articleRepo) ListArticle(ctx context.Context, g *biz.ListArticleRequest) ([]*biz.Article, error) {
	query := repo.buildQuery(g)

	list, err := query.WithAuthor().
		Limit(g.Limit).Offset(g.Offset).
		Order(ent.Desc(article.FieldID)).
		All(ctx)

	var lstArticle = make([]*biz.Article, 0, len(list))

	if err = copier.Copy(&lstArticle, list); err != nil {
		return nil, err
	}
	return lstArticle, err
}

func (repo *articleRepo) CountArticle(ctx context.Context, g *biz.ListArticleRequest) (int, error) {
	query := repo.buildQuery(g)
	return query.Count(ctx)
}

func (repo *articleRepo) DeleteArticle(ctx context.Context, slug uuid.UUID) (int, error) {
	return repo.data.db.Article.Delete().Where(article.SlugEQ(slug)).Exec(ctx)
}

func (repo *articleRepo) buildQuery(g *biz.ListArticleRequest) *ent.ArticleQuery {
	query := repo.data.db.Article.Query()

	if g.Title != "" {
		query.Where(article.TitleContains(g.Title))
	}

	if g.AuthorId != 0 {
		query.Where(article.AuthorIDEQ(g.AuthorId))
	}

	if g.Tag != "" {
		query.Where(article.TagsContains(g.Tag))
	}

	if g.UserIds != nil {
		query.Where(article.AuthorIDIn(g.UserIds...))
	}

	return query
}
