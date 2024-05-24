package auth_v1

import (
	"context"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/utils"
	"github.com/pkg/errors"
	"os"
)

func (s *serv) GetAccessToken(_ context.Context, token string) (string, error) {
	claims, err := utils.VerifyToken(token, []byte(os.Getenv("REFRESH_TOKEN_SECRET_KEY")))
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	accessToken, err := utils.GenerateToken(
		claims.Subject,
		[]byte(os.Getenv("ACCESS_TOKEN_SECRET_KEY")),
		accessTokenExpiration)
	if err != nil {
		return "", errors.New("failed to generate access token")
	}

	return accessToken, nil
}
