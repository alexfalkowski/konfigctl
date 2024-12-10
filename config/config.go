package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/konfigctl/client"
)

// Config for the client.
type Config struct {
	Client         *client.Config `yaml:"client,omitempty" json:"client,omitempty" toml:"client,omitempty"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

// Valid or error.
func (c Config) Valid() error {
	if c.Client == nil || c.Config == nil {
		return config.ErrInvalidConfig
	}

	return c.Config.Valid()
}

func decorateConfig(cfg *Config) *config.Config {
	return cfg.Config
}

func clientConfig(cfg *Config) *client.Config {
	return cfg.Client
}
