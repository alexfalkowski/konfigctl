package main

import (
	"os"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/konfigctl/cmd"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

func command() *sc.Command {
	c := sc.New(cmd.Version)
	c.RegisterInput(c.Root(), "env:KONFIG_CONFIG_FILE")

	co := c.AddClient("config", "Get Config.", cmd.ConfigOptions...)
	c.RegisterOutput(co, "env:KONFIG_APP_CONFIG_FILE")

	se := c.AddClient("secrets", "Write secrets.", cmd.SecretsOptions...)
	c.RegisterOutput(se, "env:KONFIG_APP_CONFIG_FILE")

	return c
}
