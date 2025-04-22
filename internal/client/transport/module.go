package transport

import (
	"github.com/alexfalkowski/konfigctl/internal/client/transport/grpc"
	"github.com/alexfalkowski/konfigctl/internal/client/transport/http"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	grpc.Module,
	http.Module,
	fx.Provide(NewService),
)
