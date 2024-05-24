package auth_v1

import (
	"context"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/schema"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"strconv"
	"sync"
)

func (s *serv) Register(ctx context.Context, user *schema.AuthRegister) (*schema.AuthToken, error) {
	var wg sync.WaitGroup

	newUserId, err := s.authV1Repository.Create(ctx, user)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var accessToken string
	var refreshToken string

	wg.Add(1)
	go func(subject string) {
		defer wg.Done()
		accessToken, _ = utils.GenerateToken(
			subject,
			[]byte(os.Getenv("ACCESS_TOKEN_SECRET_KEY")),
			accessTokenExpiration)
	}(strconv.FormatInt(newUserId, 10))

	wg.Add(1)
	go func(subject string) {
		defer wg.Done()
		refreshToken, _ = utils.GenerateToken(
			subject,
			[]byte(os.Getenv("REFRESH_TOKEN_SECRET_KEY")),
			refreshTokenExpiration)
	}(strconv.FormatInt(newUserId, 10))

	wg.Wait()

	return &schema.AuthToken{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
