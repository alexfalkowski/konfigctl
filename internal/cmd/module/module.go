package module

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/module"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/konfigctl/internal/client"
	"github.com/alexfalkowski/konfigctl/internal/config"
	"github.com/alexfalkowski/konfigctl/internal/token"
	"go.uber.org/fx"
)

// Module got all commands.
var Module = fx.Options(
	module.Module, token.Module, http.Module,
	telemetry.Module, client.Module,
	config.Module, cmd.Module,
)
