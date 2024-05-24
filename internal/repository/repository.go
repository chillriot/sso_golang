package repository

import (
	"context"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/schema"
)

type AuthV1Repository interface {
	GetUserByUsername(ctx context.Context, username string) (*schema.UserPassword, error)
	Create(ctx context.Context, user *schema.AuthRegister) (int64, error)
}
