package http

import (
	"context"

	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/id"
	"github.com/alexfalkowski/go-service/net/http/rpc"
	"github.com/alexfalkowski/go-service/telemetry/logger"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/alexfalkowski/go-service/token"
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/konfigctl/internal/client/config"
	v1 "github.com/alexfalkowski/konfigctl/internal/client/konfig/v1"
	"go.uber.org/fx"
)

// ClientParams for konfig.
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

// NewClient for http.
func NewClient(params ClientParams) (*rpc.Client, error) {
	client, err := http.NewClient(
		http.WithClientLogger(params.Logger), http.WithClientTracer(params.Tracer),
		http.WithClientMetrics(params.Meter), http.WithClientRetry(params.Client.Retry),
		http.WithClientUserAgent(params.UserAgent), http.WithClientTimeout(params.Client.Timeout),
		http.WithClientTokenGenerator(params.Generator), http.WithClientTLS(params.Client.TLS),
		http.WithClientID(params.ID))
	if err != nil {
		return nil, err
	}

	return rpc.NewClient(params.Client.Address,
		rpc.WithClientRoundTripper(client.Transport),
		rpc.WithClientContentType("application/protojson"),
		rpc.WithClientTimeout(params.Client.Timeout),
	), nil
}

// NewService for grpc.
func NewService(client *rpc.Client) *Service {
	return &Service{client: client}
}

// Service for grpc.
type Service struct {
	client *rpc.Client
}

// GetConfig for a specific application.
func (s *Service) GetConfig(ctx context.Context, req *v1.GetConfigRequest) (*v1.GetConfigResponse, error) {
	resp := &v1.GetConfigResponse{}
	err := s.client.Invoke(ctx, v1.Service_GetConfig_FullMethodName, req, resp)

	return resp, err
}

// GetSecrets that are configured.
func (s *Service) GetSecrets(ctx context.Context, req *v1.GetSecretsRequest) (*v1.GetSecretsResponse, error) {
	resp := &v1.GetSecretsResponse{}
	err := s.client.Invoke(ctx, v1.Service_GetSecrets_FullMethodName, req, resp)

	return resp, err
}
