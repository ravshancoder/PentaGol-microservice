package postgres

import "github.com/jmoiron/sqlx"

type AdminRepo struct {
	db *sqlx.DB
}

func NewAdminRepo(db *sqlx.DB) *AdminRepo {
	return &AdminRepo{
		db: db,
	}
}
