package repo

import (
	"context"
	"monorepo/src/auth_service/pkg/entity"
)

// Defining Base interface for Authentication
type ICustomerStorage interface {
	CustomerLogin(ctx context.Context, req entity.LoginRequest) (entity.CreateResponse, error)
	SignUp(ctx context.Context, customer entity.Customer) (entity.Customer, error)
}
