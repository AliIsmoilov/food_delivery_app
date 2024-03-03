package storage

import (
	"monorepo/src/auth_service/storage/postgres"
	"monorepo/src/auth_service/storage/repo"

	"gorm.io/gorm"
)

type IStorage interface {
	Authenitication() repo.IAuthStorage
	Customer() repo.ICustomerStorage
}

type storagePg struct {
	authRepo     repo.IAuthStorage
	customerRepo repo.ICustomerStorage
}

func New(db *gorm.DB) *storagePg {
	return &storagePg{
		authRepo:     postgres.New(db),
		customerRepo: postgres.NewCustomer(db),
	}
}

func (s storagePg) Authenitication() repo.IAuthStorage {
	return s.authRepo
}

func (s storagePg) Customer() repo.ICustomerStorage {
	return s.customerRepo
}
