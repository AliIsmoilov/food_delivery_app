package storage

import (
	"fmt"
	"monorepo/src/order_service/storage/postgres"
	"monorepo/src/order_service/storage/repo"

	"gorm.io/gorm"
)

type IStorage interface {
	Order() repo.IOderStorage
}

type storagePg struct {
	orderRepo repo.IOderStorage
}

func New(db *gorm.DB) IStorage {
	fmt.Println("storage ")
	storage := &storagePg{
		orderRepo: postgres.NewOrderRepo(db),
	}
	return storage
}

func (s storagePg) Order() repo.IOderStorage {
	return s.orderRepo
}
