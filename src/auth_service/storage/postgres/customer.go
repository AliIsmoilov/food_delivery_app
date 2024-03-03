package postgres

import (
	"context"
	"errors"
	"fmt"
	"monorepo/src/auth_service/pkg/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type customerRepo struct {
	db *gorm.DB
}

func NewCustomer(db *gorm.DB) *customerRepo {
	return &customerRepo{db: db}
}

func (r *customerRepo) CustomerLogin(ctx context.Context, req entity.LoginRequest) (entity.CreateResponse, error) {

	customer := entity.Customer{}
	res := r.db.
		Table("customers").
		Where("phone_number = ?", req.PhoneNumber).
		Where("deleted_at IS NULL").
		First(&customer)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return entity.CreateResponse{IsExist: false}, nil
		}
		return entity.CreateResponse{}, res.Error
	}

	var deviceExist bool
	result := r.db.Raw("SELECT exists(SELECT * FROM customers_device WHERE customer_id=? AND device_id=?)", customer.ID, req.DeviceId).
		Scan(&deviceExist)
	if result.Error != nil {
		return entity.CreateResponse{}, fmt.Errorf("error in Login(2): %w", result.Error)
	}

	if deviceExist {
		result = r.db.Exec("UPDATE customers_device SET device_name=?, notification_id=?, updated_at=current_timestamp WHERE customer_id=? AND device_id=?",
			req.DeviceName, req.NotificationId, customer.ID, req.DeviceId)
		if result.Error != nil {
			return entity.CreateResponse{}, fmt.Errorf("error in Login(3): %w", result.Error)
		}
	} else {

		customerDevice := entity.DeviceInfo{
			Id:             uuid.New(),
			CustomerId:     uuid.MustParse(customer.ID.String()),
			DeviceId:       req.DeviceInfo.DeviceId,
			DeviceName:     req.DeviceInfo.DeviceName,
			NotificationId: req.DeviceInfo.NotificationId,
		}

		res := r.db.Table("customers_device").Create(&customerDevice)
		if res.Error != nil {
			return entity.CreateResponse{}, res.Error
		}
	}

	return entity.CreateResponse{
		ID:      customer.ID.String(),
		Name:    customer.Name,
		Photo:   customer.Photo,
		Email:   customer.Email,
		IsExist: true,
	}, nil
}

func (r *customerRepo) SignUp(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	// Start a transaction
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Insert the new customer record using GORM
	if err := tx.Table("customers").Create(&customer).Error; err != nil {
		tx.Rollback()
		return entity.Customer{}, fmt.Errorf("error in SignUp(1): %w", err)
	}

	if err := tx.Table("customers_device").Create(&customer.DeviceInfo).Error; err != nil {
		tx.Rollback()
		return entity.Customer{}, fmt.Errorf("error in SignUp(2): %w", err)
	}

	// Commit the transaction if successful
	if err := tx.Commit().Error; err != nil {
		return entity.Customer{}, fmt.Errorf("error in SignUp(3): %w", err)
	}

	return customer, nil
}
