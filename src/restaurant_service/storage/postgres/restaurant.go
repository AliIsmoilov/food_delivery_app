package postgres

import (
	"context"
	"fmt"
	"monorepo/src/restaurant_service/entity"
	"monorepo/src/restaurant_service/pkg/db/constants"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type restaurantRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *restaurantRepo {
	return &restaurantRepo{db: db}
}

func (r *restaurantRepo) CreateRestaurant(ctx context.Context, restaurant entity.Restaurant) (uuid.UUID, error) {

	res := r.db.WithContext(ctx).Create(&restaurant)
	if res.Error != nil {
		return uuid.UUID{}, fmt.Errorf("error in CreateRestaurant: %w", res.Error)
	} else if res.RowsAffected == 0 {
		return uuid.UUID{}, fmt.Errorf("error in CreateRestaurant: %w", constants.ErrRowsAffectedZero)
	}

	return restaurant.Id, nil
}

func (r *restaurantRepo) ListRestaurants(ctx context.Context, restaurant entity.ListRestaurantsReq) (entity.ListRestaurantsResp, error) {
	restaurants := []entity.Restaurant{}

	res := r.db.WithContext(ctx).
		Table("restaurants").
		Where("deleted_at IS NULL").
		Find(&restaurants)
	if res.Error != nil {
		return entity.ListRestaurantsResp{}, fmt.Errorf("error in ListRestaurants: %w", res.Error)
	}

	return entity.ListRestaurantsResp{
		Count:       len(restaurants),
		Restaurants: restaurants,
	}, nil
}
