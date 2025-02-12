package client

import (
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(NewServiceClient),
	fx.Provide(NewClient),
)
