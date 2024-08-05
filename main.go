package main

import (
	"os"

	"github.com/alexfalkowski/go-client-template/cmd"
	sc "github.com/alexfalkowski/go-service/cmd"
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
