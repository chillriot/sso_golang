package auth_v1

import (
	"context"
	descAuthV1 "gitea.24example.ru/RosarStoreBackend/protobuf/pkg/sso_v1"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/converter"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (i *Implementation) Register(ctx context.Context, req *descAuthV1.RegisterRequest) (*descAuthV1.AuthTokenResponse, error) {
	token, err := i.authV1Service.Register(ctx, converter.AuthRegisterFromProtoToService(req))

	if err != nil {
		log.Printf("register error: %v\n", err)
		return nil, status.Errorf(codes.InvalidArgument, "register failed: invalid argument")
	}

	return converter.AuthLoginFromServiceToProto(token), nil
}
