package cmd

import (
	"github.com/alexfalkowski/go-service/module"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/token"
	"github.com/alexfalkowski/konfigctl/client"
	"github.com/alexfalkowski/konfigctl/cmd/secrets"
	"github.com/alexfalkowski/konfigctl/config"
	"go.uber.org/fx"
)

// SecretsOptions for cmd.
var SecretsOptions = []fx.Option{
	module.Module, token.Module,
	telemetry.Module, client.Module,
	secrets.Module, config.Module, Module,
}
