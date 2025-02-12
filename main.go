package main

import (
	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/konfigctl/internal/cmd"
)

func main() {
	command().ExitOnError()
}

func command() *sc.Command {
	command := sc.New(cmd.Version)
	command.RegisterInput(command.Root(), "env:KONFIG_CONFIG_FILE")

	cmd.RegisterConfig(command)
	cmd.RegisterSecrets(command)

	return command
}
