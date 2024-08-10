package client

import (
	"context"

	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/token"
	"github.com/alexfalkowski/go-service/transport/grpc"
	v1 "github.com/alexfalkowski/konfigctl/client/konfig/v1"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ServiceClientParams for konfig.
type ServiceClientParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Client    *Config
	Logger    *zap.Logger
	Tracer    trace.Tracer
	Meter     metric.Meter
	Generator token.Generator
	UserAgent env.UserAgent
}

// NewServiceClient for konfig.
func NewServiceClient(params ServiceClientParams) (v1.ServiceClient, error) {
	sec, err := grpc.WithClientTLS(params.Client.TLS)
	if err != nil {
		return nil, err
	}

	opts := []grpc.ClientOption{
		grpc.WithClientLogger(params.Logger), grpc.WithClientTracer(params.Tracer),
		grpc.WithClientMetrics(params.Meter), grpc.WithClientRetry(params.Client.Retry),
		grpc.WithClientUserAgent(params.UserAgent), grpc.WithClientTimeout(params.Client.Timeout),
		grpc.WithClientTokenGenerator(params.Generator), sec,
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

// Config from konfig.
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

// Secrets from konfig.
func (c *Client) Secrets(ctx context.Context) (map[string][]byte, error) {
	req := &v1.GetSecretsRequest{Secrets: c.config.Secrets.Files}

	resp, err := c.client.GetSecrets(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetSecrets(), nil
}
