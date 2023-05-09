package storage

import (
	"github.com/PentaGol/admin_service/storage/postgres"
	"github.com/PentaGol/admin_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Admin() repo.AdminStoreI
}

type storagePg struct {
	db       *sqlx.DB
	adminRepo repo.AdminStoreI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		adminRepo: postgres.NewAdminRepo(db),
	}
}

func (s storagePg) Admin() repo.AdminStoreI {
	return s.adminRepo
}
