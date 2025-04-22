package transport

import (
	"context"

	"github.com/alexfalkowski/konfigctl/internal/client/config"
	v1 "github.com/alexfalkowski/konfigctl/internal/client/konfig/v1"
	"github.com/alexfalkowski/konfigctl/internal/client/transport/grpc"
	"github.com/alexfalkowski/konfigctl/internal/client/transport/http"
)

// NewService returns the service according to the correct transport.
func NewService(conf *config.Client, grpc *grpc.Service, http *http.Service) Service {
	if conf.IsHTTP() {
		return http
	}

	return grpc
}

// Service for different transports.
type Service interface {
	// GetConfig for a specific application.
	GetConfig(ctx context.Context, req *v1.GetConfigRequest) (*v1.GetConfigResponse, error)

	// GetSecrets that are configured.
	GetSecrets(ctx context.Context, req *v1.GetSecretsRequest) (*v1.GetSecretsResponse, error)
}
