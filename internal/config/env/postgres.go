package env

import (
	"errors"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/config"
	"os"
)

var _ config.PGConfig = (*pgConfig)(nil)

const (
	host     = "POSTGRES_HOST"
	db       = "POSTGRES_DB"
	user     = "POSTGRES_USER"
	password = "POSTGRES_PASSWORD"
	port     = "POSTGRES_PORT"
)

type pgConfig struct {
	dsn string
}

func NewPGConfig() (*pgConfig, error) {
	dsn := "host=" + os.Getenv(host) +
		" port=" + os.Getenv(port) +
		" dbname=" + os.Getenv(db) +
		" user=" + os.Getenv(user) +
		" password=" + os.Getenv(password) +
		" sslmode=disable"

	if len(dsn) == 0 {
		return nil, errors.New("pg dsn not found")
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
