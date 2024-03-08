package storage

import (
	"fmt"
	"monorepo/src/order_service/storage/postgres"
	"monorepo/src/order_service/storage/repo"

	"gorm.io/gorm"
)

type IStorage interface {
	Order() repo.IOderStorage
	Courier() repo.ICourierStorage
}

type storagePg struct {
	orderRepo   repo.IOderStorage
	courierRepo repo.ICourierStorage
}

func New(db *gorm.DB) IStorage {
	fmt.Println("storage ")
	storage := &storagePg{
		orderRepo:   postgres.NewOrderRepo(db),
		courierRepo: postgres.NewCourierRepo(db),
	}
	return storage
}

func (s storagePg) Order() repo.IOderStorage {
	return s.orderRepo
}

func (s storagePg) Courier() repo.ICourierStorage {
	return s.courierRepo
}
