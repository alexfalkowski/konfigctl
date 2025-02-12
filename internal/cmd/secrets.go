package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/module"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/konfigctl/internal/client"
	"github.com/alexfalkowski/konfigctl/internal/cmd/secrets"
	"github.com/alexfalkowski/konfigctl/internal/config"
	"github.com/alexfalkowski/konfigctl/internal/token"
)

// RegisterSecrets for cmd.
func RegisterSecrets(command *cmd.Command) {
	client := command.AddClient("secrets", "Write secrets.",
		module.Module, token.Module,
		telemetry.Module, client.Module,
		secrets.Module, config.Module, Module,
	)
	command.RegisterOutput(client, "env:KONFIG_APP_CONFIG_FILE")
}
