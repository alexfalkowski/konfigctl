package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/id"
	"github.com/alexfalkowski/go-service/telemetry/logger"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/alexfalkowski/go-service/token"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/konfigctl/internal/client/config"
	v1 "github.com/alexfalkowski/konfigctl/internal/client/konfig/v1"
	"go.uber.org/fx"
)

// ClientParams for grpc.
type ClientParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Tracer    *tracer.Tracer
	Meter     *metrics.Meter
	ID        id.Generator
	Client    *config.Client
	Logger    *logger.Logger
	Generator token.Generator
	UserAgent env.UserAgent
}

// NewServiceClient for grpc.
func NewServiceClient(params ClientParams) (v1.ServiceClient, error) {
	conn, err := grpc.NewClient(params.Client.Address,
		grpc.WithClientLogger(params.Logger), grpc.WithClientTracer(params.Tracer),
		grpc.WithClientMetrics(params.Meter), grpc.WithClientRetry(params.Client.Retry),
		grpc.WithClientUserAgent(params.UserAgent), grpc.WithClientTimeout(params.Client.Timeout),
		grpc.WithClientTokenGenerator(params.Generator), grpc.WithClientTLS(params.Client.TLS),
		grpc.WithClientID(params.ID),
	)

	params.Lifecycle.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			return conn.Close()
		},
	})

	return v1.NewServiceClient(conn), err
}

// NewService for grpc.
func NewService(client v1.ServiceClient) *Service {
	return &Service{client: client}
}

// Service for grpc.
type Service struct {
	client v1.ServiceClient
}

// GetConfig for a specific application.
func (s *Service) GetConfig(ctx context.Context, req *v1.GetConfigRequest) (*v1.GetConfigResponse, error) {
	return s.client.GetConfig(ctx, req)
}

// GetSecrets that are configured.
func (s *Service) GetSecrets(ctx context.Context, req *v1.GetSecretsRequest) (*v1.GetSecretsResponse, error) {
	return s.client.GetSecrets(ctx, req)
}
