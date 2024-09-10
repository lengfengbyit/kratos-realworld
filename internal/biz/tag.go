package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type TagRepo interface {
	ListTag(ctx context.Context, limit, offset int) ([]string, error)
}

type TagUsecase struct {
	repo TagRepo
	log  *log.Helper
}

func NewTagUsecase(repo TagRepo, logger log.Logger) *TagUsecase {
	return &TagUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (biz *TagUsecase) ListTag(ctx context.Context, limit, offset int) ([]string, error) {
	return biz.repo.ListTag(ctx, limit, offset)
}
