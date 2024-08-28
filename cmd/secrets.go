package cmd

import (
	"github.com/alexfalkowski/go-service/compress"
	"github.com/alexfalkowski/go-service/encoding"
	"github.com/alexfalkowski/go-service/sync"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/konfigctl/client"
	"github.com/alexfalkowski/konfigctl/cmd/secrets"
	"github.com/alexfalkowski/konfigctl/config"
	"go.uber.org/fx"
)

// SecretsOptions for cmd.
var SecretsOptions = []fx.Option{
	sync.Module, compress.Module, encoding.Module,
	telemetry.Module, client.Module,
	secrets.Module, config.Module, Module,
}
