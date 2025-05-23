package main

import (
	"context"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/konfigctl/internal/cmd"
)

func main() {
	command().ExitOnError(context.Background())
}

func command() *sc.Command {
	command := sc.New(env.NewName(), env.NewVersion())

	cmd.RegisterConfig(command)
	cmd.RegisterSecrets(command)

	return command
}
