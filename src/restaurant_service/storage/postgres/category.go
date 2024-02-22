package postgres

import (
	"context"
	"fmt"
	"monorepo/src/restaurant_service/entity"
	"monorepo/src/restaurant_service/pkg/db/constants"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *categoryRepo {
	return &categoryRepo{db: db}
}

func (r *categoryRepo) CreateCategory(ctx context.Context, category entity.Category) (uuid.UUID, error) {

	res := r.db.WithContext(ctx).Create(&category)
	if res.Error != nil {
		return uuid.UUID{}, fmt.Errorf("error in CreateCategory: %w", res.Error)
	} else if res.RowsAffected == 0 {
		return uuid.UUID{}, fmt.Errorf("error in CreateCategory: %w", constants.ErrRowsAffectedZero)
	}

	return category.Id, nil
}

func (r *categoryRepo) ListCategories(ctx context.Context, restaurant entity.ListCategoriesReq) (entity.ListCategoriesResp, error) {
	categories := []entity.Category{}

	res := r.db.WithContext(ctx).
		Table("categories").
		Where("deleted_at IS NULL").
		Find(&categories)
	if res.Error != nil {
		return entity.ListCategoriesResp{}, fmt.Errorf("error in ListCategories: %w", res.Error)
	}

	return entity.ListCategoriesResp{
		Count:      len(categories),
		Categories: categories,
	}, nil
}
