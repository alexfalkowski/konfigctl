package client

import (
	"github.com/alexfalkowski/konfigctl/internal/client/transport"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	transport.Module,
	fx.Provide(NewClient),
)
