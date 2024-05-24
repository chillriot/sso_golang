package auth_v1

import (
	"gitea.24example.ru/RosarStoreBackend/protobuf/pkg/sso_v1"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/service"
)

type Implementation struct {
	sso_v1.UnimplementedAuthServiceServer
	sso_v1.UnimplementedTokenServiceServer
	authV1Service service.AuthV1Service
}

func NewImplementation(authV1Service service.AuthV1Service) *Implementation {
	return &Implementation{authV1Service: authV1Service}
}
