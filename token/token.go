package token

import (
	"context"

	"github.com/alexfalkowski/go-service/os"
	"github.com/alexfalkowski/go-service/token"
)

// NewGenerator for konfig.
func NewGenerator(cfg *Config) token.Generator {
	return &Generator{cfg: cfg}
}

// Generator for konfig.
type Generator struct {
	cfg *Config
}

// Generate token from secret file.
func (g *Generator) Generate(ctx context.Context) (context.Context, []byte, error) {
	d, err := os.ReadBase64File(g.cfg.Key)

	return ctx, []byte(d), err
}
