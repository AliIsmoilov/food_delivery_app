package postgres

import (
	"context"
	"errors"
	"fmt"
	"monorepo/src/order_service/entity"
	"monorepo/src/order_service/pkg/db/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type courierRepo struct {
	db *gorm.DB
}

func NewCourierRepo(db *gorm.DB) *courierRepo {
	return &courierRepo{db: db}
}

func (r *courierRepo) CreateCourier(ctx context.Context, courier entity.Courier) (uuid.UUID, error) {

	res := r.db.WithContext(ctx).Create(&courier)
	if res.Error != nil {
		return uuid.UUID{}, fmt.Errorf("error in CreateCourier: %w", res.Error)
	} else if res.RowsAffected == 0 {
		return uuid.UUID{}, fmt.Errorf("error in CreateCourier: %w", constants.ErrRowsAffectedZero)
	}

	return courier.ID, nil
}

func (r *courierRepo) CourierList(ctx context.Context, req entity.CourierListReq) (entity.CourierListResp, error) {
	couriers := []entity.Courier{}
	var count int64

	res := r.db.WithContext(ctx).
		Table("couriers").
		Where("deleted_at IS NULL")

	if req.Search != "" {
		req.Search = "%" + req.Search + "%"
		res.Where("name ilike ?", req.Search)
	}

	tx := res.Count(&count)
	if tx.Error != nil {
		return entity.CourierListResp{}, res.Error
	}

	if req.Page > 0 && req.Limit > 0 {
		offset := (req.Page - 1) * req.Limit
		res = res.Offset(int(offset)).
			Limit(int(req.Limit))
	} else if req.Limit > 0 {
		res = res.Limit(int(req.Limit))
	}

	res = res.Find(&couriers)
	if res.Error != nil {
		return entity.CourierListResp{}, fmt.Errorf("error in CourierList: %w", res.Error)
	}

	return entity.CourierListResp{
		Count:    len(couriers),
		Couriers: couriers,
	}, nil
}

func (w *courierRepo) UpdateCourier(ctx context.Context, req entity.Courier) error {
	res := w.db.Model(entity.Courier{}).
		Where("id=?", req.ID).
		Updates(req)
	if res.Error != nil {
		return fmt.Errorf("error in UpdateCourier: %w", res.Error)
	} else if res.RowsAffected == 0 {
		return fmt.Errorf("error in UpdateCourier: %w", constants.ErrRowsAffectedZero)
	}

	return nil
}

func (wr *courierRepo) DeleteCourier(ctx context.Context, id string) error {
	res := wr.db.Table("couriers").
		Where("id=?", id).
		Update("deleted_at", time.Now())
	if res.Error != nil {
		return fmt.Errorf("error in DeleteCourier: %w", res.Error)
	} else if res.RowsAffected == 0 {
		return fmt.Errorf("error in DeleteCourier: %w", constants.ErrRowsAffectedZero)
	}

	return nil
}

func (wr *courierRepo) GetCourier(ctx context.Context, id string) (entity.Courier, error) {
	resp := entity.Courier{}
	res := wr.db.
		Table("couriers").
		Where("id = ?", id).
		First(&resp)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return entity.Courier{}, fmt.Errorf("courier not found %w", constants.ErrRecordNotFound)
		}
		return entity.Courier{}, res.Error
	}

	return resp, nil
}
