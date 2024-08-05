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
	cl := c.AddClient(cmd.ClientOptions...)
	c.RegisterInput(cl, "")
	c.RegisterOutput(cl, "")

	return c
}
