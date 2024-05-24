package auth_v1

import (
	"context"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/utils"
	"github.com/pkg/errors"
	"os"
)

func (s *serv) GetRefreshToken(_ context.Context, token string) (string, error) {
	claims, err := utils.VerifyToken(token, []byte(os.Getenv("REFRESH_TOKEN_SECRET_KEY")))
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	refreshToken, err := utils.GenerateToken(
		claims.Subject,
		[]byte(os.Getenv("REFRESH_TOKEN_SECRET_KEY")),
		refreshTokenExpiration)
	if err != nil {
		return "", errors.New("failed to generate refresh token")
	}

	return refreshToken, nil
}
