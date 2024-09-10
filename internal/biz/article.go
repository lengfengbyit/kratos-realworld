package biz

import (
	"context"
	"kratos-realworld/internal/data/ent"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

var (
	ErrInvalidSlug = errors.New(401, "params error", "invalid slug")
)

// Article VO
type Article struct {
	Id             int64
	AuthorId       int64
	Slug           string
	Title          string
	Description    string
	Body           string
	TagList        []string
	Favorited      bool
	FavoritesCount int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Author         *Author
}
type Author struct {
	ID       int64
	Username string
	Image    string
	Bio      string
	Followed bool
}

type ListArticleRequest struct {
	AuthorId int64
	Tag      string
	Title    string
	UserIds  []int64

	Limit  int
	Offset int
}

// Tags 使用 copier 从 ent.Article 转换时使用
// 从 string 类型的 Tags 转换为 Article.TagList
func (a *Article) Tags(tags string) {
	a.TagList = strings.Split(tags, ",")
}

func (a *Article) Edges(edges ent.ArticleEdges) {
	if edges.Author != nil {
		var author Author
		if err := copier.Copy(&author, edges.Author); err != nil {
			return
		}
		a.Author = &author
	}
}

// ArticleRepo interface, 定义 data 层的接口
type ArticleRepo interface {
	CreateArticle(ctx context.Context, g *Article) (*Article, error)
	UpdateArticle(ctx context.Context, g *Article) (int, error)
	GetArticle(ctx context.Context, slug uuid.UUID) (*Article, error)
	ListArticle(ctx context.Context, g *ListArticleRequest) ([]*Article, error)
	DeleteArticle(ctx context.Context, slug uuid.UUID) (int, error)
	CountArticle(ctx context.Context, g *ListArticleRequest) (int, error)
}

// ArticleUsecase 业务层的操作
type ArticleUsecase struct {
	repo ArticleRepo
	log  *log.Helper
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (a *ArticleUsecase) CreateArticle(ctx context.Context, g *Article) (*Article, error) {
	return a.repo.CreateArticle(ctx, g)
}

func (a *ArticleUsecase) UpdateArticle(ctx context.Context, g *Article) (int, error) {
	return a.repo.UpdateArticle(ctx, g)
}

func (a *ArticleUsecase) GetArticle(ctx context.Context, slug string) (*Article, error) {
	slugUUID := a.slugToUUID(slug)
	if slugUUID == uuid.Nil {
		return nil, ErrInvalidSlug
	}

	return a.repo.GetArticle(ctx, slugUUID)
}

func (a *ArticleUsecase) ListArticle(ctx context.Context, g *ListArticleRequest) (int, []*Article, error) {
	total, err := a.repo.CountArticle(ctx, g)
	if err != nil {
		return 0, nil, errors.InternalServer("query article count", err.Error())
	}

	if total == 0 {
		return 0, []*Article{}, nil
	}

	list, err := a.repo.ListArticle(ctx, g)
	return total, list, err
}

func (a *ArticleUsecase) DeleteArticle(ctx context.Context, slug string) (int, error) {
	slugUUID := a.slugToUUID(slug)
	if slugUUID == uuid.Nil {
		return 0, ErrInvalidSlug
	}

	return a.repo.DeleteArticle(ctx, slugUUID)
}

func (a *ArticleUsecase) slugToUUID(slug string) uuid.UUID {
	slugUUID, err := uuid.Parse(slug)
	if err != nil {
		a.log.Errorf("slug invalid, err: %s", err)
		return uuid.Nil
	}
	return slugUUID
}
