package auth_v1

import (
	"context"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/client/db"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/repository"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/repository/auth_v1/model"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/schema"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/utils"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"strings"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.AuthV1Repository {
	return &repo{db: db}
}

func (r *repo) GetUserByUsername(ctx context.Context, username string) (*schema.UserPassword, error) {
	username = strings.ToLower(username)

	builder := sq.Select("e24up.user_id, e24up.password").
		From("public.e24_users e24u").
		Join("public.e24_users_password e24up on e24u.id = e24up.user_id").
		Where(sq.Or{
			sq.Eq{"e24u.email": username},
			sq.Eq{"e24u.phone": username},
		}).
		PlaceholderFormat(sq.Dollar).
		OrderBy("e24up.created_at DESC").
		Limit(1)

	query, args, _ := builder.ToSql()

	q := db.Query{
		Name:     "auth_v1_repository.GetUserByUsername",
		QueryRaw: query,
	}

	var user model.UserPassword
	err := r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &schema.UserPassword{
		UserID:   user.UserId,
		Password: user.Password,
	}, nil
}

func (r *repo) Create(ctx context.Context, user *schema.AuthRegister) (int64, error) {
	if len(user.Password) == 0 {
		return 0, errors.New("password is empty")
	}
	hashPassword, _ := utils.HashPassword(user.Password)

	email := strings.ToLower(user.Email)
	phone := strings.ToLower(user.Phone)

	builder := sq.
		Insert("e24_users").
		PlaceholderFormat(sq.Dollar).
		Columns("email", "fullname", "phone, is_active, is_admin").
		Values(email, user.Fullname, phone, true, false).
		Suffix("RETURNING id")

	query, args, _ := builder.ToSql()

	q := db.Query{
		Name:     "auth_v1_repository.Create",
		QueryRaw: query,
	}

	var id int64

	if err := r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		if strings.Contains(err.Error(), "23505") {
			return 0, errors.New("User already exists")
		}
		return 0, err
	}

	builder = sq.Insert("e24_users_password").
		PlaceholderFormat(sq.Dollar).
		Columns("user_id", "password").
		Values(id, hashPassword)
	query, args, _ = builder.ToSql()

	q = db.Query{
		Name:     "auth_v1_repository.CreatePassword",
		QueryRaw: query,
	}

	r.db.DB().QueryRowContext(ctx, q, args...)

	return id, nil
}
