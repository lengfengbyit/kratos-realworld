package biz

import (
	"context"
	"kratos-realworld/internal/data/ent"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type Comment struct {
	Id        int64
	ArticleId int64
	UserId    int64
	Body      string
	CreatedAt time.Time
	Author    *Author
}

func (a *Comment) Edges(edges ent.CommentEdges) {
	if edges.Author != nil {
		var author Author
		if err := copier.Copy(&author, edges.Author); err != nil {
			return
		}
		a.Author = &author
	}
}

type CommentRepo interface {
	CreateComment(ctx context.Context, g *Comment) (*Comment, error)
	DeleteComment(ctx context.Context, id int64) error
	ListComment(ctx context.Context, articleId int64, limit, offset int) ([]*Comment, error)
}

type CommonUsecase struct {
	repo CommentRepo
	log  *log.Helper
}

func NewCommonUsecase(repo CommentRepo, logger log.Logger) *CommonUsecase {
	return &CommonUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (biz *CommonUsecase) CreateComment(ctx context.Context, g *Comment) (*Comment, error) {
	return biz.repo.CreateComment(ctx, g)
}

func (biz *CommonUsecase) DeleteComment(ctx context.Context, id int64) error {
	return biz.repo.DeleteComment(ctx, id)
}

func (biz *CommonUsecase) ListComment(ctx context.Context, articleId int64, limit, offset int) ([]*Comment, error) {
	return biz.repo.ListComment(ctx, articleId, limit, offset)
}
