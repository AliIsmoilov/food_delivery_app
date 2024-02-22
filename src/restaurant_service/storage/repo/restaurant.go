package repo

import (
	"context"
	"monorepo/src/restaurant_service/entity"

	"github.com/google/uuid"
)

// Defining Base interface for Authentication
type IResraurantStorage interface {
	CreateRestaurant(ctx context.Context, restaurant entity.Restaurant) (uuid.UUID, error)
	ListRestaurants(ctx context.Context, restaurant entity.ListRestaurantsReq) (entity.ListRestaurantsResp, error)
}
