package data

import (
	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type profileRepo struct {
	data *Data
	log  *log.Helper
}

// NewProfileRepo .
func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
