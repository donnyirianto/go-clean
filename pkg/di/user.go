package di

import (
	"github.com/google/wire"

	"github.com/donnyirianto/go-clean/pkg/api/handler"
	"github.com/donnyirianto/go-clean/pkg/repository"
	"github.com/donnyirianto/go-clean/pkg/usecase"
)

var UserSet = wire.NewSet(
	repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler,
)
