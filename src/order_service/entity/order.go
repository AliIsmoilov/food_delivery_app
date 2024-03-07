package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id                uuid.UUID `gorm:"id"`
	Code              string    `gorm:"code"`
	Note              string    `gorm:"note"`
	TotalPrice        uint32    `json:"total_price"`
	Status            string    `gorm:"default:new"`
	RestaurantId      string    `gorm:"restaurant_id"`
	CustomerId        string    `gorm:"customer_id"`
	CourierId         string    `gorm:"courier_id"`
	Address           string    `gorm:"address"`
	Lat               float64   `gorm:"lat"`
	Long              float64   `gorm:"long"`
	Items             []Item
	DeliveredAt       *time.Time `gorm:"<-:update column:delivered_at"`
	DeliveryStartedAt *time.Time `gorm:"<-:update column:delivery_started_at"`
	ReadyAt           *time.Time `gorm:"<-:update column:ready_at"`
	CreatedAt         *time.Time `gorm:"column:created_at"`
	UpdatedAt         *time.Time `gorm:"<-:update column:updated_at"`
	DeletedAt         *time.Time `gorm:"<-:update column:deleted_at"`
}

type Item struct {
	ID        string     `gorm:"id"`
	ItemID    string     `gorm:"item_id"`
	OrderID   string     `gorm:"column:order_id" json:"-"`
	Qty       uint32     `gorm:"qty"`
	Price     uint32     `gorm:"price"`
	Order     Order      `gorm:"foreignKey:OrderID" json:"-"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"<-:update column:updated_at"`
	DeletedAt *time.Time `gorm:"<-:update column:deleted_at"`
}

type OrderListReq struct {
	Status string
}

type OrderListResp struct {
	Count  int
	Orders []Order
}
