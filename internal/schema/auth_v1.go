package schema

import (
	"database/sql"
	"github.com/golang-jwt/jwt"
	"time"
)

type User struct {
	ID int64

	Info UserInfo

	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserInfo struct {
	Fullname  sql.NullString
	Email     string
	Phone     string
	PathImage sql.NullString
	BirthDate sql.NullTime
	IsActive  bool
	IsAdmin   bool
}

type UserPassword struct {
	UserID   int64
	Password string
}

type AuthLogin struct {
	Username string
	Password string
}

type AuthRegister struct {
	Fullname string
	Password string
	Email    string
	Phone    string
}

type AuthToken struct {
	AccessToken  string
	RefreshToken string
}

type AccessToken struct {
	AccessToken string
}

type RefreshToken struct {
	RefreshToken string
}

type UserClaims struct {
	jwt.StandardClaims
}
