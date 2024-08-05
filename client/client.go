package client

import (
	"context"

	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/security/token"
	v1 "github.com/alexfalkowski/konfigctl/client/konfig/v1"
	"github.com/alexfalkowski/konfigctl/transport/grpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// NewGenerator for token.
func NewGenerator(tkn token.Tokenizer) token.Generator {
	return tkn
}

// ServiceClientParams for gRPC.
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

// NewServiceClient for gRPC.
func NewServiceClient(params ServiceClientParams) (v1.ServiceClient, error) {
	opts := grpc.ClientOpts{
		Lifecycle: params.Lifecycle, Client: params.Client.Config,
		Logger: params.Logger, Tracer: params.Tracer, Meter: params.Meter,
		Generator: params.Generator, UserAgent: params.UserAgent,
	}
	conn, err := grpc.NewClient(opts)

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

// Config from client.
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

// Secrets from client.
func (c *Client) Secrets(ctx context.Context) (map[string][]byte, error) {
	req := &v1.GetSecretsRequest{Secrets: c.config.Secrets.Files}

	resp, err := c.client.GetSecrets(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetSecrets(), nil
}
