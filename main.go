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

	co := command.AddClient("config", "Get Config.", cmd.ConfigOptions...)
	command.RegisterOutput(co, "env:KONFIG_APP_CONFIG_FILE")

	se := command.AddClient("secrets", "Write secrets.", cmd.SecretsOptions...)
	command.RegisterOutput(se, "env:KONFIG_APP_CONFIG_FILE")

	return command
}
