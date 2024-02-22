package models

import (
	"fmt"
	libsUtils "monorepo/src/libs/utils"
)

type Restaurant struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Adress      string  `json:"adress"`
	PhoneNumber string  `json:"phone_number"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	Photo       string  `json:"photo"`
}

// Validate StuffLoginModel
func (r *Restaurant) Validate() error {

	if !libsUtils.IsPhoneValid(r.PhoneNumber) {
		return fmt.Errorf("%s is not valid phone", r.PhoneNumber)
	}

	return nil
}

type Response struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Item struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	CategoryId   string `json:"category_id"`
	RestaurantId string `json:"restaurant_id"`
	Description  string `json:"description"`
	Photo        string `json:"photo"`
	Price        int32  `json:"price"`
}
