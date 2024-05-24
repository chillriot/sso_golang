package env

import (
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/config"
	"github.com/pkg/errors"
	"net"
	"os"
)

var _ config.GRPCConfig = (*gRPCConfig)(nil)

type gRPCConfig struct {
	host string
	port string
}

func NewGRPCConfig() (*gRPCConfig, error) {
	host := os.Getenv("GRPC_HOST")

	if len(host) == 0 {
		return nil, errors.New("GRPC_HOST env var required")
	}

	port := os.Getenv("GRPC_PORT")
	if len(port) == 0 {
		return nil, errors.New("GRPC_PORT env var required")
	}

	cfg := &gRPCConfig{
		host: host,
		port: port,
	}
	return cfg, nil
}

func (cfg *gRPCConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
