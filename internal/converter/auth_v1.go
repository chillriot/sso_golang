package converter

import (
	descAuthV1 "gitea.24example.ru/RosarStoreBackend/protobuf/pkg/sso_v1"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/schema"
)

func AuthLoginFromServiceToProto(token *schema.AuthToken) *descAuthV1.AuthTokenResponse {
	return &descAuthV1.AuthTokenResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}
}

func AccessTokenFromServiceToProto(accessToken *schema.AccessToken) *descAuthV1.GetAccessTokenResponse {
	return &descAuthV1.GetAccessTokenResponse{
		AccessToken: accessToken.AccessToken,
	}
}

func RefreshTokenFromServiceToProto(refreshToken *schema.RefreshToken) *descAuthV1.GetRefreshTokenResponse {
	return &descAuthV1.GetRefreshTokenResponse{
		RefreshToken: refreshToken.RefreshToken,
	}
}

func AuthLoginFromProtoToService(auth *descAuthV1.LoginRequest) *schema.AuthLogin {
	return &schema.AuthLogin{
		Username: auth.Username,
		Password: auth.Password,
	}
}

func AuthRegisterFromProtoToService(auth *descAuthV1.RegisterRequest) *schema.AuthRegister {
	return &schema.AuthRegister{
		Fullname: auth.Fullname,
		Password: auth.Password,
		Email:    auth.Email,
		Phone:    auth.Phone,
	}
}

func RefreshTokenFromProtoToService(token *descAuthV1.GetRefreshTokenRequest) *schema.RefreshToken {
	return &schema.RefreshToken{
		RefreshToken: token.RefreshToken,
	}
}
