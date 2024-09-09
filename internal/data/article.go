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

	var art biz.Article
	if err = copier.Copy(&art, info); err != nil {
		return nil, err
	}
	return &art, err
}

func (repo *articleRepo) ListArticle(ctx context.Context, g *biz.Article, tag string, limit, offset int) ([]*biz.Article, error) {
	query := repo.buildQuery(tag, g)

	list, err := query.WithAuthor().
		Limit(limit).Offset(offset).
		Order(ent.Desc(article.FieldID)).
		All(ctx)

	var lstArticle = make([]*biz.Article, 0, len(list))

	if err = copier.Copy(&lstArticle, list); err != nil {
		return nil, err
	}
	return lstArticle, err
}

func (repo *articleRepo) CountArticle(ctx context.Context, g *biz.Article, tag string) (int, error) {
	query := repo.buildQuery(tag, g)
	return query.Count(ctx)
}

func (repo *articleRepo) DeleteArticle(ctx context.Context, slug uuid.UUID) (int, error) {
	return repo.data.db.Article.Delete().Where(article.SlugEQ(slug)).Exec(ctx)
}

func (repo *articleRepo) buildQuery(tag string, g *biz.Article) *ent.ArticleQuery {
	query := repo.data.db.Article.Query()

	if g.Title != "" {
		query.Where(article.TitleContains(g.Title))
	}

	if g.AuthorId != 0 {
		query.Where(article.AuthorIDEQ(g.AuthorId))
	}

	if tag != "" {
		query.Where(article.TagsContains(tag))
	}
	return query
}
