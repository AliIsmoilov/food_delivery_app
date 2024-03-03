package entity

import (
	"time"

	"github.com/google/uuid"
)

type DeviceInfo struct {
	Id             uuid.UUID `gorm:"primaryKey"`
	DeviceId       string
	CustomerId     uuid.UUID
	DeviceName     string
	NotificationId string
	CreatedAt      time.Time
	UpdatedAt      time.Time `gorm:"<-:update"`
	DeletedAt      time.Time `gorm:"<-:update"`
}

type LoginRequest struct {
	PhoneNumber string
	DeviceInfo
}

type CreateResponse struct {
	ID      string
	Name    string
	Email   string
	Photo   string
	IsExist bool
}

// Customer ...
type Customer struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	PhoneNumber string
	Name        string
	Email       string
	Photo       string
	DeviceInfo  DeviceInfo `gorm:"-"`
}
