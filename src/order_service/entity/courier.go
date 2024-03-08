package entity

import (
	"time"

	"github.com/google/uuid"
)

type Courier struct {
	ID          uuid.UUID
	Name        string
	PhoneNumber string
	Photo       string
	CarModel    string
	CarNumber   string
	CarColor    string
	IsAvailable bool
	Latitude    float32
	Longitude   float32
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"<-:update column:updated_at"`
	DeletedAt   *time.Time `gorm:"<-:update column:deleted_at"`
}

type CourierListReq struct {
	Page   uint32
	Limit  uint32
	Search string
}

type CourierListResp struct {
	Count    int
	Couriers []Courier
}
