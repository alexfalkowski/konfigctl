package client

import (
	"context"

	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/id"
	"github.com/alexfalkowski/go-service/telemetry/logger"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/alexfalkowski/go-service/token"
	"github.com/alexfalkowski/go-service/transport/grpc"
	v1 "github.com/alexfalkowski/konfigctl/internal/client/konfig/v1"
	"go.uber.org/fx"
)

// ServiceClientParams for konfig.
type ServiceClientParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Tracer    *tracer.Tracer
	Meter     *metrics.Meter
	ID        id.Generator
	Client    *Config
	Logger    *logger.Logger
	Generator token.Generator
	UserAgent env.UserAgent
}

// NewServiceClient for konfig.
func NewServiceClient(params ServiceClientParams) (v1.ServiceClient, error) {
	opts := []grpc.ClientOption{
		grpc.WithClientLogger(params.Logger), grpc.WithClientTracer(params.Tracer),
		grpc.WithClientMetrics(params.Meter), grpc.WithClientRetry(params.Client.Retry),
		grpc.WithClientUserAgent(params.UserAgent), grpc.WithClientTimeout(params.Client.Timeout),
		grpc.WithClientTokenGenerator(params.Generator), grpc.WithClientTLS(params.Client.TLS),
		grpc.WithClientID(params.ID),
	}
	conn, err := grpc.NewClient(params.Client.Address, opts...)

	params.Lifecycle.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			return conn.Close()
		},
	})

	return v1.NewServiceClient(conn), err
}

// Client for konfig.
type Client struct {
	client v1.ServiceClient
	config *Config
}

// NewClient for konfig.
func NewClient(client v1.ServiceClient, config *Config) *Client {
	return &Client{client: client, config: config}
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

	resp, err := c.client.GetConfig(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetConfig().GetData(), nil
}

// Secrets from request.
func (c *Client) Secrets(ctx context.Context) (map[string][]byte, error) {
	req := &v1.GetSecretsRequest{Secrets: c.config.Secrets.Files}

	resp, err := c.client.GetSecrets(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetSecrets(), nil
}
