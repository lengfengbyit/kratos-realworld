package data

import (
	"context"
	"kratos-realworld/internal/biz"
	"kratos-realworld/internal/data/ent/comment"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type commentRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &commentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c commentRepo) CreateComment(ctx context.Context, g *biz.Comment) (*biz.Comment, error) {
	info, err := c.data.db.Comment.Create().
		SetUserID(g.UserId).
		SetArticleID(g.ArticleId).
		SetBody(g.Body).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	var comm biz.Comment
	if err = copier.Copy(&comm, info); err != nil {
		return nil, err
	}

	return &comm, err
}

func (c commentRepo) DeleteComment(ctx context.Context, id int64) error {
	return c.data.db.Comment.DeleteOneID(id).Exec(ctx)
}

func (c commentRepo) ListComment(ctx context.Context, articleId int64, limit, offset int) ([]*biz.Comment, error) {
	lst, err := c.data.db.Comment.Query().
		WithAuthor().
		Where(comment.ArticleID(articleId)).
		Offset(offset).Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var comments = make([]*biz.Comment, 0, len(lst))
	if err = copier.Copy(&comments, lst); err != nil {
		return nil, err
	}
	return comments, nil
}
