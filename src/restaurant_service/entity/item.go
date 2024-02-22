package entity

import (
	"github.com/google/uuid"
)

type Item struct {
	Id           uuid.UUID `gorm:"id"`
	Name         string    `gorm:"name"`
	CategoryId   string    `gorm:"category_id"`
	Category     Category  `gorm:"foreignKey:category_id"`
	RestaurantId string    `gorm:"restaurant_id"`
	Description  string    `gorm:"description"`
	Photo        string    `gorm:"photo"`
	Price        int32     `gorm:"price"`
}

type RestaurantItemsGroupedByCategoryReq struct {
	RestaurantId uuid.UUID
}

type RestaurantItemsGroupedByCategoryModel struct {
	Category Category
	Items    []Item
}

type RestaurantItemsGroupedByCategoryResp struct {
	Items []RestaurantItemsGroupedByCategoryModel
}
