package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID int64 `db:"id"`

	Info UserInfo `db:""`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type UserInfo struct {
	Fullname  string       `db:"fullname"`
	Email     string       `db:"email"`
	Phone     string       `db:"phone"`
	PathImage string       `db:"path_image"`
	BirthDate sql.NullTime `db:"birthday"`
	IsActive  bool         `db:"is_active"`
	IsAdmin   bool         `db:"is_admin"`
}

type UserPassword struct {
	UserId   int64  `db:"user_id"`
	Password string `db:"password"`
}
