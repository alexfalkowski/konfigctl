package config

import (
	"github.com/alexfalkowski/go-service/config"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(NewConfig),
	config.Module,
	fx.Decorate(decorateConfig),
	fx.Provide(clientConfig),
	fx.Provide(tokenConfig),
)
