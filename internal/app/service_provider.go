package app

import (
	"context"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/api/auth_v1"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/client/db"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/client/db/postgres"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/closer"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/config"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/config/env"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/repository"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/service"
	"log"

	authV1Repository "gitea.24example.ru/RosarStoreBackend/sso_v1/internal/repository/auth_v1"
	authV1Service "gitea.24example.ru/RosarStoreBackend/sso_v1/internal/service/auth_v1"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	gRPCConfig config.GRPCConfig
	httpConfig config.HTTPConfig

	dbClient         db.Client
	authV1Repository repository.AuthV1Repository

	authV1Service service.AuthV1Service

	authV1Impl *auth_v1.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to load postgres config: %v", err)
		}
		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to load http config: %v", err)
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.gRPCConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to load gRPC config: %v", err)
		}
		s.gRPCConfig = cfg
	}
	return s.gRPCConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		client, err := postgres.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client %v", err)
		}

		if err = client.DB().Ping(ctx); err != nil {
			log.Fatalf("failed to ping db: %v", err)
		}

		closer.Add(client.Close)
		s.dbClient = client
	}

	return s.dbClient
}

func (s *serviceProvider) AuthV1Repository(ctx context.Context) repository.AuthV1Repository {
	if s.authV1Repository == nil {
		s.authV1Repository = authV1Repository.NewRepository(s.DBClient(ctx))
	}

	return s.authV1Repository
}

func (s *serviceProvider) AuthV1Service(ctx context.Context) service.AuthV1Service {
	if s.authV1Service == nil {
		s.authV1Service = authV1Service.NewService(s.AuthV1Repository(ctx))
	}

	return s.authV1Service
}

func (s *serviceProvider) AuthV1Impl(ctx context.Context) *auth_v1.Implementation {
	if s.authV1Impl == nil {
		s.authV1Impl = auth_v1.NewImplementation(s.AuthV1Service(ctx))
	}

	return s.authV1Impl
}
