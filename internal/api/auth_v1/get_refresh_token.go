package auth_v1

import (
	"context"
	descAuthV1 "gitea.24example.ru/RosarStoreBackend/protobuf/pkg/sso_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (i *Implementation) GetRefreshToken(ctx context.Context, req *descAuthV1.GetRefreshTokenRequest) (*descAuthV1.GetRefreshTokenResponse, error) {
	token, err := i.authV1Service.GetRefreshToken(ctx, req.RefreshToken)

	if err != nil {
		log.Printf("Get refresh token error: %v\n", err)
		return nil, status.Errorf(codes.Internal, "Get refresh token: Internal")
	}

	return &descAuthV1.GetRefreshTokenResponse{RefreshToken: token}, nil
}
