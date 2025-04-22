package client

import (
	"context"

	"github.com/alexfalkowski/konfigctl/internal/client/config"
	v1 "github.com/alexfalkowski/konfigctl/internal/client/konfig/v1"
	"github.com/alexfalkowski/konfigctl/internal/client/transport"
)

// Client for konfig.
type Client struct {
	config  *config.Client
	service transport.Service
}

// NewClient for konfig.
func NewClient(config *config.Client, service transport.Service) *Client {
	return &Client{config: config, service: service}
}

// Config from request.
func (c *Client) Config(ctx context.Context) ([]byte, error) {
	cfg := c.config.Configuration
	req := &v1.GetConfigRequest{
		Application: cfg.Application,
		Version:     cfg.Version,
		Environment: cfg.Environment,
		Continent:   cfg.Continent,
		Country:     cfg.Country,
		Command:     cfg.Command,
		Kind:        cfg.Kind,
	}

	resp, err := c.service.GetConfig(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetConfig().GetData(), nil
}

// Secrets from request.
func (c *Client) Secrets(ctx context.Context) (map[string][]byte, error) {
	req := &v1.GetSecretsRequest{Secrets: c.config.Secrets.Files}

	resp, err := c.service.GetSecrets(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetSecrets(), nil
}
