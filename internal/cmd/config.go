package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/module"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/konfigctl/internal/client"
	"github.com/alexfalkowski/konfigctl/internal/cmd/config"
	kc "github.com/alexfalkowski/konfigctl/internal/config"
	"github.com/alexfalkowski/konfigctl/internal/token"
)

// RegisterConfig for cmd.
func RegisterConfig(command *cmd.Command) {
	client := command.AddClient("config", "Get Config.", module.Module, token.Module,
		telemetry.Module, config.Module,
		client.Module, kc.Module,
		config.Module, Module,
	)
	command.RegisterOutput(client, "env:KONFIG_APP_CONFIG_FILE")
}
