package auth_v1

import (
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/repository"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/service"
)

type serv struct {
	authV1Repository repository.AuthV1Repository
}

func NewService(authV1Repository repository.AuthV1Repository) service.AuthV1Service {
	return &serv{
		authV1Repository: authV1Repository,
	}
}
