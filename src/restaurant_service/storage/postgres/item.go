package postgres

import (
	"context"
	"fmt"
	"monorepo/src/restaurant_service/entity"
	"monorepo/src/restaurant_service/pkg/db/constants"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) *itemRepo {
	return &itemRepo{db: db}
}

func (r *itemRepo) CreateItem(ctx context.Context, item entity.Item) (uuid.UUID, error) {

	res := r.db.WithContext(ctx).Create(&item)
	if res.Error != nil {
		return uuid.UUID{}, fmt.Errorf("error in CreateItem: %w", res.Error)
	} else if res.RowsAffected == 0 {
		return uuid.UUID{}, fmt.Errorf("error in CreateItem: %w", constants.ErrRowsAffectedZero)
	}

	return item.Id, nil
}

func (r *itemRepo) RestaurantItemsGroupedByCategory(ctx context.Context, restaurant entity.RestaurantItemsGroupedByCategoryReq) (entity.RestaurantItemsGroupedByCategoryResp, error) {
	items := []entity.Item{}

	res := r.db.WithContext(ctx).
		Table("items").
		Where("deleted_at IS NULL").
		Preload("Category").
		Find(&items)
	if res.Error != nil {
		return entity.RestaurantItemsGroupedByCategoryResp{}, fmt.Errorf("error in RestaurantItemsGroupedByCategory: %w", res.Error)
	}

	itemsMap := make(map[entity.Category][]entity.Item)
	for _, item := range items {
		// _, ok := itemsMap[item.CategoryId]
		itemsMap[item.Category] = append(itemsMap[item.Category], item)
	}

	resp := entity.RestaurantItemsGroupedByCategoryResp{}
	for category, items := range itemsMap {
		resp.Items = append(resp.Items, entity.RestaurantItemsGroupedByCategoryModel{
			Category: category,
			Items:    items,
		})
	}

	return resp, nil
}
