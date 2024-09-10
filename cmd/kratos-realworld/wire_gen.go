// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-realworld/internal/biz"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/data"
	"kratos-realworld/internal/server"
	"kratos-realworld/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, logger)
	userApiService := service.NewUserApiService(userUsecase, confServer)
	grpcServer := server.NewGRPCServer(confServer, greeterService, userApiService, logger)
	followRepo := data.NewFollowRepo(dataData, logger)
	profileUsecase := biz.NewProfileUsecase(userRepo, followRepo, logger)
	profileService := service.NewProfileService(profileUsecase)
	articleRepo := data.NewArticleRepo(dataData, logger)
	articleUsecase := biz.NewArticleUsecase(articleRepo, logger)
	favoriteRepo := data.NewFavoriteRepo(dataData, logger)
	favoriteUsecase := biz.NewFavoriteUsecase(favoriteRepo, logger)
	articleService := service.NewArticleService(articleUsecase, profileUsecase, favoriteUsecase, logger)
	commentRepo := data.NewCommentRepo(dataData, logger)
	commonUsecase := biz.NewCommonUsecase(commentRepo, logger)
	commentService := service.NewCommentService(commonUsecase, articleUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, greeterService, userApiService, profileService, articleService, commentService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
