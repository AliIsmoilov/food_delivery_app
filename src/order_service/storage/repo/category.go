package repo

import (
	"context"
	"monorepo/src/restaurant_service/entity"

	"github.com/google/uuid"
)

// Defining Base interface for Authentication
type ICategoryStorage interface {
	CreateCategory(ctx context.Context, category entity.Category) (uuid.UUID, error)
	ListCategories(ctx context.Context, restaurant entity.ListCategoriesReq) (entity.ListCategoriesResp, error)
}
