package service

import (
	"context"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/schema"
)

type AuthV1Service interface {
	Login(ctx context.Context, user *schema.AuthLogin) (*schema.AuthToken, error)
	Register(ctx context.Context, user *schema.AuthRegister) (*schema.AuthToken, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
}
