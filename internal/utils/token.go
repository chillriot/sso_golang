package utils

import (
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/schema"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"time"
)

func GenerateToken(subject string, secretKey []byte, duration time.Duration) (string, error) {
	claims := schema.UserClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   subject,
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func VerifyToken(tokenStr string, secretKey []byte) (*schema.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&schema.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.Errorf("unexpected token signing method")
			}

			return secretKey, nil
		},
	)
	if err != nil {
		return nil, errors.Errorf("invalid token: %s", err.Error())
	}

	claims, ok := token.Claims.(*schema.UserClaims)
	if !ok {
		return nil, errors.Errorf("invalid token claims")
	}

	return claims, nil
}
