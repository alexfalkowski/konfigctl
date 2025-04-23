package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/konfigctl/internal/cmd/module"
	"github.com/alexfalkowski/konfigctl/internal/cmd/secrets"
)

// RegisterSecrets for cmd.
func RegisterSecrets(command *cmd.Command) {
	flags := command.AddClient("secrets", "Write secrets.", module.Module, secrets.Module)
	flags.AddInput("")
}
