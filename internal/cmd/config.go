package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/konfigctl/internal/cmd/config"
	"github.com/alexfalkowski/konfigctl/internal/cmd/module"
)

// RegisterConfig for cmd.
func RegisterConfig(command *cmd.Command) {
	flags := command.AddClient("config", "Get Config.", module.Module, config.Module)
	flags.AddInput("")
	flags.AddOutput("")
}
