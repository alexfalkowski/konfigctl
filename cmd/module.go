package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/konfigctl/token"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	token.Module,
	cmd.Module,
	env.Module,
	fx.Provide(NewVersion),
)
