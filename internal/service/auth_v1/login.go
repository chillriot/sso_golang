package auth_v1

import (
	"context"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/schema"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/utils"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	refreshTokenExpiration = 360 * time.Minute
	accessTokenExpiration  = 60 * time.Minute
)

func (s *serv) Login(ctx context.Context, user *schema.AuthLogin) (*schema.AuthToken, error) {
	var wg sync.WaitGroup

	userPassword, err := s.authV1Repository.GetUserByUsername(ctx, user.Username)
	if err != nil {
		return nil, err
	}

	if !utils.VerifyPassword(userPassword.Password, user.Password) {
		return nil, errors.New("invalid password")
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
	}(strconv.FormatInt(userPassword.UserID, 10))

	wg.Add(1)
	go func(subject string) {
		defer wg.Done()
		refreshToken, _ = utils.GenerateToken(
			subject,
			[]byte(os.Getenv("REFRESH_TOKEN_SECRET_KEY")),
			refreshTokenExpiration)
	}(strconv.FormatInt(userPassword.UserID, 10))

	wg.Wait()
	return &schema.AuthToken{AccessToken: accessToken, RefreshToken: refreshToken}, nil

}
