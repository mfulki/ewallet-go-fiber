package repository

import "database/sql"

type UserRepository interface {
}

type UserRepositoryPostgres struct {
	db *sql.DB
}