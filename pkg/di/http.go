package di

import (
	"github.com/google/wire"

	api "github.com/donnyirianto/go-clean/pkg/api"
	middleware "github.com/donnyirianto/go-clean/pkg/api/middleware"
)

var HTTPSet = wire.NewSet(
	api.NewServerHTTP,
	middleware.NewErrorHandler,
	middleware.NewAuthentication,
	wire.Struct(new(api.Middlewares), "*"),
	wire.Struct(new(api.Handlers), "*"),
)
