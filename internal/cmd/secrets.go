package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/module"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/konfigctl/internal/client"
	"github.com/alexfalkowski/konfigctl/internal/cmd/secrets"
	"github.com/alexfalkowski/konfigctl/internal/config"
	"github.com/alexfalkowski/konfigctl/internal/token"
)

// RegisterSecrets for cmd.
func RegisterSecrets(command *cmd.Command) {
	flags := flags.NewFlagSet("secrets")

	command.RegisterInput(flags, "env:KONFIG_CONFIG_FILE")
	command.RegisterOutput(flags, "env:KONFIG_APP_CONFIG_FILE")
	command.AddClient("secrets", "Write secrets.", flags,
		module.Module, token.Module,
		telemetry.Module, client.Module,
		secrets.Module, config.Module, cmd.Module,
	)
}
