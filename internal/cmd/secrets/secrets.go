package secrets

import (
	"context"
	"io/fs"
	"path/filepath"

	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/os"
	"github.com/alexfalkowski/konfigctl/internal/client"
	"github.com/alexfalkowski/konfigctl/internal/client/config"
	"go.uber.org/fx"
)

// Params for secrets.
type Params struct {
	fx.In

	Lifecycle    fx.Lifecycle
	Client       *client.Client
	OutputConfig *cmd.OutputConfig
	Config       *config.Client
	FileSystem   os.FileSystem
}

// Start for secrets.
func Start(params Params) {
	cmd.Start(params.Lifecycle, func(ctx context.Context) error {
		secrets, err := params.Client.Secrets(ctx)
		if err != nil {
			return err
		}

		cfg := params.Config.Secrets

		for name, secret := range secrets {
			path := filepath.Join(cfg.Path, name)

			err := params.FileSystem.WriteFile(path, secret, fs.FileMode(cfg.Mode))
			if err != nil {
				return err
			}
		}

		return nil
	})
}
