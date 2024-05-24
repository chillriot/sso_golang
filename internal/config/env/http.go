package env

import (
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/config"
	"github.com/pkg/errors"
	"net"
	"os"
)

var _ config.HTTPConfig = (*httpConfig)(nil)

type httpConfig struct {
	host string
	port string
}

func NewHTTPConfig() (*httpConfig, error) {
	host := os.Getenv("HTTP_HOST")

	if len(host) == 0 {
		return nil, errors.New("HTTP_HOST env var required")
	}

	port := os.Getenv("HTTP_PORT")
	if len(port) == 0 {
		return nil, errors.New("HTTP_PORT env var required")
	}

	cfg := &httpConfig{
		host: host,
		port: port,
	}
	return cfg, nil
}

func (cfg *httpConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
