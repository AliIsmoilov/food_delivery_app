package storage

import (
	"fmt"
	"monorepo/src/restaurant_service/storage/postgres"
	"monorepo/src/restaurant_service/storage/repo"

	"gorm.io/gorm"
)

type IStorage interface {
	Restaurant() repo.IResraurantStorage
	Category() repo.ICategoryStorage
	Item() repo.IItemStorage
}

type storagePg struct {
	restaurantRepo repo.IResraurantStorage
	categoryRepo   repo.ICategoryStorage
	itemRepo       repo.IItemStorage
}

func New(db *gorm.DB) IStorage {
	fmt.Println("storage ")
	storage := &storagePg{
		restaurantRepo: postgres.New(db),
		categoryRepo:   postgres.NewCategoryRepo(db),
		itemRepo:       postgres.NewItemRepo(db),
	}
	return storage
}

func (s storagePg) Restaurant() repo.IResraurantStorage {
	return s.restaurantRepo
}

func (s storagePg) Category() repo.ICategoryStorage {
	return s.categoryRepo
}

func (s storagePg) Item() repo.IItemStorage {
	return s.itemRepo
}
