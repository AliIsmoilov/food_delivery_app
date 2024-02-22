package repo

import (
	"context"
	"monorepo/src/restaurant_service/entity"

	"github.com/google/uuid"
)

type IItemStorage interface {
	CreateItem(ctx context.Context, item entity.Item) (uuid.UUID, error)
	RestaurantItemsGroupedByCategory(ctx context.Context, restaurant entity.RestaurantItemsGroupedByCategoryReq) (entity.RestaurantItemsGroupedByCategoryResp, error)
}
