package auth_v1

import (
	"context"
	descAuthV1 "gitea.24example.ru/RosarStoreBackend/protobuf/pkg/sso_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (i *Implementation) GetAccessToken(ctx context.Context, req *descAuthV1.GetAccessTokenRequest) (*descAuthV1.GetAccessTokenResponse, error) {
	token, err := i.authV1Service.GetAccessToken(ctx, req.RefreshToken)

	if err != nil {
		log.Printf("Get access token error: %v\n", err)
		return nil, status.Errorf(codes.Internal, "Get access token: Internal")
	}

	return &descAuthV1.GetAccessTokenResponse{AccessToken: token}, nil
}
