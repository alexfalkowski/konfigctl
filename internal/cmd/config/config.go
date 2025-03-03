package config

import (
	"context"
	"io/fs"

	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/konfigctl/internal/client"
	"go.uber.org/fx"
)

// Params for config.
type Params struct {
	fx.In

	Lifecycle    fx.Lifecycle
	Client       *client.Client
	OutputConfig *cmd.OutputConfig
	Config       *client.Config
}

// Start for config.
func Start(params Params) {
	cmd.Start(params.Lifecycle, func(ctx context.Context) {
		config, err := params.Client.Config(ctx)
		runtime.Must(err)

		err = params.OutputConfig.Write(config, fs.FileMode(params.Config.Configuration.Mode))
		runtime.Must(err)
	})
}
