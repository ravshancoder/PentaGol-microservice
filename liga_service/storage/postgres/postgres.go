package postgres

import "github.com/jmoiron/sqlx"

type Repo struct {
	db *sqlx.DB
}

func NewLigaRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func NewGameRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func NewClubRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}