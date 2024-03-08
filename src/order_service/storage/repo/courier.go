package repo

import (
	"context"
	"monorepo/src/order_service/entity"

	"github.com/google/uuid"
)

// Defining Base interface for Authentication
type ICourierStorage interface {
	CreateCourier(ctx context.Context, courier entity.Courier) (uuid.UUID, error)
	CourierList(ctx context.Context, req entity.CourierListReq) (entity.CourierListResp, error)
	UpdateCourier(ctx context.Context, req entity.Courier) error
	DeleteCourier(ctx context.Context, id string) error
	GetCourier(ctx context.Context, id string) (entity.Courier, error)
}
