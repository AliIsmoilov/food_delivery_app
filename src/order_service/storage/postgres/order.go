package postgres

import (
	"context"
	"fmt"
	"monorepo/src/order_service/entity"
	"monorepo/src/order_service/pkg/db/constants"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *orderRepo {
	return &orderRepo{db: db}
}

func (r *orderRepo) CreateOrder(ctx context.Context, order entity.Order) (uuid.UUID, error) {

	res := r.db.WithContext(ctx).Create(&order)
	if res.Error != nil {
		return uuid.UUID{}, fmt.Errorf("error in CreateOrder: %w", res.Error)
	} else if res.RowsAffected == 0 {
		return uuid.UUID{}, fmt.Errorf("error in CreateOrder: %w", constants.ErrRowsAffectedZero)
	}

	return order.Id, nil
}

func (r *orderRepo) OrderList(ctx context.Context, req entity.OrderListReq) (entity.OrderListResp, error) {
	orders := []entity.Order{}

	res := r.db.WithContext(ctx).
		Table("orders").
		Where("deleted_at IS NULL")

	if req.Status != "" {
		res.Where("status = ?", req.Status)
	}

	res = res.Find(&orders)
	if res.Error != nil {
		return entity.OrderListResp{}, fmt.Errorf("error in OrderList: %w", res.Error)
	}

	return entity.OrderListResp{
		Count:  len(orders),
		Orders: orders,
	}, nil
}
