package repository

import (
	"context"
	"database/sql"

	"github.com/mfulki/ewallet-go-fiber/apperrors"
	"github.com/mfulki/ewallet-go-fiber/db"
	"github.com/mfulki/ewallet-go-fiber/entity"
)

type PasswordTokenRepository interface {
	InsertToken(ctx context.Context, userId int, token string) (*entity.TokenPassword, error)
	DeleteToken(ctx context.Context, userId int) error
	CheckToken(ctx context.Context, token string) (*int, error)
}

type passwordTokenRepositoryImpl struct {
	db db.DBConnection
}

func NewPasswordTokenRepository(db db.DBConnection) *passwordTokenRepositoryImpl {
	return &passwordTokenRepositoryImpl{
		db: db,
	}
}
func (r *passwordTokenRepositoryImpl) InsertToken(ctx context.Context, userId int, token string) (*entity.TokenPassword, error) {
	q := `insert into password_tokens (user_id,token) values ($1,$2) returning expired_at`
	t := entity.TokenPassword{}
	err := r.db.QueryRowContext(ctx, q, userId, token).Scan(&t.ExpiredAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound()
		}
		return nil, apperrors.ErrInternalServer()
	}
	t.Token = token
	return &t, nil
}
func (r *passwordTokenRepositoryImpl) DeleteToken(ctx context.Context, userId int) error {
	q := `update password_tokens set deleted_at =now() where user_id =$1 and deleted_at is null; `
	_, err := r.db.ExecContext(ctx, q, userId)
	if err != nil {
		return apperrors.ErrInternalServer()
	}
	return nil
}

func (r *passwordTokenRepositoryImpl) CheckToken(ctx context.Context, token string) (*int, error) {
	var userId *int
	q := `select user_id from password_tokens where token=$1 and deleted_at is null and now()<expired_at`
	err := r.db.QueryRowContext(ctx, q, token).Scan(&userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrForbidden()
		}
		return nil, apperrors.ErrInternalServer()
	}
	return userId, nil
}
