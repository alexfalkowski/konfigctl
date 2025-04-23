package config

import (
	"github.com/alexfalkowski/go-service/config"
	client "github.com/alexfalkowski/konfigctl/internal/client/config"
)

// Config for the client.
type Config struct {
	Client         *client.Client `yaml:"client,omitempty" json:"client,omitempty" toml:"client,omitempty"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

func decorateConfig(cfg *Config) *config.Config {
	return cfg.Config
}

func clientConfig(cfg *Config) *client.Client {
	return cfg.Client
}
