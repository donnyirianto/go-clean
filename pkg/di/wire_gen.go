// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/donnyirianto/go-clean/pkg/api"
	"github.com/donnyirianto/go-clean/pkg/api/handler"
	"github.com/donnyirianto/go-clean/pkg/api/middleware"
	"github.com/donnyirianto/go-clean/pkg/config"
	"github.com/donnyirianto/go-clean/pkg/driver/db"
	"github.com/donnyirianto/go-clean/pkg/driver/log/adapter"
	config2 "github.com/donnyirianto/go-clean/pkg/driver/log/config"
	"github.com/donnyirianto/go-clean/pkg/repository"
	"github.com/donnyirianto/go-clean/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {
	errorHandler := middleware.NewErrorHandler()
	authentication := middleware.NewAuthentication()
	middlewares := &api.Middlewares{
		ErrorHandler:   errorHandler,
		Authentication: authentication,
	}
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	handlers := api.Handlers{
		UserHandler: userHandler,
	}
	logger, err := config2.ProvidZapLogger()
	if err != nil {
		return nil, err
	}
	zapImplement := adapter.ProvideLogger(logger)
	serverHTTP := api.NewServerHTTP(middlewares, handlers, zapImplement, cfg)
	return serverHTTP, nil
}
