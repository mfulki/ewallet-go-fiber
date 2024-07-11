package repository

import (
	"context"
	"database/sql"

	"github.com/mfulki/ewallet-go-fiber/apperrors"
	"github.com/mfulki/ewallet-go-fiber/db"
	"github.com/mfulki/ewallet-go-fiber/entity"
	query "github.com/mfulki/ewallet-go-fiber/sql"
)

type UserRepository interface {
	GetRegisteredUserIdByEmail(Email string, ctx context.Context) (*int, error)
	GetRegisteredUserId(username string, ctx context.Context) (*int, error)
	RegisterUser(user *entity.User, ctx context.Context) (*entity.User, error)
	Login(user *entity.User, ctx context.Context) (*entity.User, error)
	GetUserDetails(ctx context.Context, username string) (*entity.User, error)
	GetUserIdFromEmail(ctx context.Context, user *entity.User) (*entity.User, error)
	ResetPassword(ctx context.Context, password string, userId int) error
}

type userRepositoryImpl struct {
	db db.DBConnection
}

func NewUserRepository(db db.DBConnection) *userRepositoryImpl {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) GetRegisteredUserId(username string, ctx context.Context) (*int, error) {
	q := query.CheckUserRegistered
	var registeredId *int
	err := r.db.QueryRowContext(ctx, q, username).Scan(&registeredId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound()
		}
		return nil, apperrors.ErrInternalServer()
	}
	return registeredId, err

}

func (r *userRepositoryImpl) GetRegisteredUserIdByEmail(email string, ctx context.Context) (*int, error) {
	q := query.CheckUserRegistered
	var registeredId *int
	err := r.db.QueryRowContext(ctx, q, email).Scan(&registeredId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound()
		}
		return nil, apperrors.ErrInternalServer()
	}
	return registeredId, err

}
func (r *userRepositoryImpl) RegisterUser(user *entity.User, ctx context.Context) (*entity.User, error) {
	q := query.Register
	err := r.db.QueryRowContext(ctx, q, user.Email, user.Password, user.FullName).Scan(&user.Id, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound()
		}
		return nil, apperrors.ErrInternalServer()
	}
	return user, err
}
func (r *userRepositoryImpl) Login(user *entity.User, ctx context.Context) (*entity.User, error) {
	q := query.Login
	err := r.db.QueryRowContext(ctx, q, user.Email).Scan(&user.Password, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound()
		}
		return nil, apperrors.ErrInternalServer()
	}
	return user, err
}

func (r *userRepositoryImpl) GetUserDetails(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	q := `select id, email from users where name=$1`
	err := r.db.QueryRowContext(ctx, q, username).Scan(&user.Id, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound()
		}
		return nil, apperrors.ErrInternalServer()
	}
	user.Name = username
	return &user, err
}
func (r *userRepositoryImpl) GetUserIdFromEmail(ctx context.Context, user *entity.User) (*entity.User, error) {
	q := `select id from users where email=$1`
	err := r.db.QueryRowContext(ctx, q, user.Email).Scan(&user.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound()
		}
		return nil, apperrors.ErrInternalServer()
	}

	return user, err
}
func (r *userRepositoryImpl) ResetPassword(ctx context.Context, password string, userId int) error {
	q := `update users set security_word=$1,updated_at=now() where id=$2`
	_, err := r.db.ExecContext(ctx, q, password, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return apperrors.ErrNotFound()
		}
		return apperrors.ErrInternalServer()
	}
	return nil
}
