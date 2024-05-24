package env

import (
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/config"
	"os"
)

var _ config.TokenConfig = (*TokenConfig)(nil)

type TokenConfig struct {
	accessTokenSecretKey  string
	refreshTokenSecretKey string
}

func (t *TokenConfig) AccessTokenSecretKey() string {
	return os.Getenv("ACCESS_TOKEN_SECRET_KEY")
}
func (t *TokenConfig) RefreshTokenSecretKey() string {
	return os.Getenv("REFRESH_TOKEN_SECRET_KEY")
}
