package entity

import "github.com/google/uuid"

type Restaurant struct {
	Id          uuid.UUID `gorm:"id"`
	Name        string    `gorm:"name"`
	Adress      string    `gorm:"adress"`
	PhoneNumber string    `gorm:"phone_number"`
	Latitude    float64   `gorm:"latitude"`
	Longitude   float64   `gorm:"longitude"`
	Photo       string    `gorm:"photo"`
}

type ListRestaurantsReq struct {
}

type ListRestaurantsResp struct {
	Count       int
	Restaurants []Restaurant
}
