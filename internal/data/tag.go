package data

import (
	"context"
	"kratos-realworld/internal/biz"
	"kratos-realworld/internal/data/ent/article"
	"maps"
	"slices"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
)

type tag struct {
	data *Data
	log  *log.Helper
}

func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &tag{data: data, log: log.NewHelper(logger)}
}

func (t *tag) ListTag(ctx context.Context, limit, offset int) ([]string, error) {
	lst, err := t.data.db.Article.Query().
		Select(article.FieldTags).
		Offset(offset).Limit(limit).
		All(ctx)

	if err != nil {
		return nil, err
	}

	var tags = make(map[string]struct{}, len(lst))
	for _, item := range lst {
		ts := strings.Split(item.Tags, ",")
		for _, k := range ts {
			tags[k] = struct{}{}
		}
	}

	return slices.Collect(maps.Keys(tags)), nil
}
