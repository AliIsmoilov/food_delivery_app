package storage

import (
	"fmt"
	"monorepo/src/order_service/storage/postgres"
	"monorepo/src/order_service/storage/repo"

	"gorm.io/gorm"
)

type IStorage interface {
	Category() repo.ICategoryStorage
}

type storagePg struct {
	categoryRepo repo.ICategoryStorage
}

func New(db *gorm.DB) IStorage {
	fmt.Println("storage ")
	storage := &storagePg{
		categoryRepo: postgres.NewCategoryRepo(db),
	}
	return storage
}

func (s storagePg) Category() repo.ICategoryStorage {
	return s.categoryRepo
}
