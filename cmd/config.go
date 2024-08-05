package cmd

import (
	"github.com/alexfalkowski/go-service/compress"
	"github.com/alexfalkowski/go-service/encoding"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/konfigctl/client"
	"github.com/alexfalkowski/konfigctl/cmd/config"
	kc "github.com/alexfalkowski/konfigctl/config"
	"go.uber.org/fx"
)

// ConfigOptions for cmd.
var ConfigOptions = []fx.Option{
	compress.Module, encoding.Module,
	telemetry.Module, config.Module,
	client.Module, kc.Module,
	config.Module, Module,
}
