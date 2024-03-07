package repo

import (
	"context"
	"monorepo/src/order_service/entity"

	"github.com/google/uuid"
)

// Defining Base interface for Authentication
type IOderStorage interface {
	CreateOrder(ctx context.Context, order entity.Order) (uuid.UUID, error)
	OrderList(ctx context.Context, req entity.OrderListReq) (entity.OrderListResp, error)
}
