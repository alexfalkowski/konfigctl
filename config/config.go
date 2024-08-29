package config

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/konfigctl/client"
	"github.com/alexfalkowski/konfigctl/token"
)

// NewConfig for config.
func NewConfig(i *cmd.InputConfig) (*Config, error) {
	c := &Config{}

	return c, i.Decode(c)
}

// IsEnabled for config.
func IsEnabled(cfg *Config) bool {
	return cfg != nil
}

// Config for the client.
type Config struct {
	Client         *client.Config `yaml:"client,omitempty" json:"client,omitempty" toml:"client,omitempty"`
	Token          *token.Config  `yaml:"token,omitempty" json:"token,omitempty" toml:"token,omitempty"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

func decorateConfig(cfg *Config) *config.Config {
	if !IsEnabled(cfg) {
		return nil
	}

	return cfg.Config
}

func clientConfig(cfg *Config) *client.Config {
	if !IsEnabled(cfg) || !client.IsEnabled(cfg.Client) {
		return nil
	}

	return cfg.Client
}

func tokenConfig(cfg *Config) *token.Config {
	if !IsEnabled(cfg) || !token.IsEnabled(cfg.Token) {
		return nil
	}

	return cfg.Token
}
