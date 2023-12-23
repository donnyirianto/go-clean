package di

import (
	"github.com/google/wire"

	log "github.com/donnyirianto/go-clean/pkg/driver/log"
	logAdapter "github.com/donnyirianto/go-clean/pkg/driver/log/adapter"
	logConfig "github.com/donnyirianto/go-clean/pkg/driver/log/config"
)

var LogSet = wire.NewSet(
	logConfig.ProvidZapLogger,
	wire.Bind(new(log.Logger),
		new(*logAdapter.ZapImplement)),
	logAdapter.ProvideLogger,
)
