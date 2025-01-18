package cmd

import (
	"github.com/alexfalkowski/go-service/module"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/konfigctl/client"
	"github.com/alexfalkowski/konfigctl/cmd/config"
	kc "github.com/alexfalkowski/konfigctl/config"
	"github.com/alexfalkowski/konfigctl/token"
	"go.uber.org/fx"
)

// ConfigOptions for cmd.
var ConfigOptions = []fx.Option{
	module.Module, token.Module,
	telemetry.Module, config.Module,
	client.Module, kc.Module,
	config.Module, Module,
}
