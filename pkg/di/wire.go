//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/donnyirianto/go-clean/pkg/api"
	"github.com/donnyirianto/go-clean/pkg/config"
	"github.com/donnyirianto/go-clean/pkg/driver/db"
)

func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, UserSet, LogSet, HTTPSet)

	return &api.ServerHTTP{}, nil
}
