package models

type Order struct {
	Id           string      `json:"id"`
	Code         string      `json:"code"`
	Note         string      `json:"note"`
	TotalPrice   uint32      `json:"total_price"`
	Status       string      `json:"status"`
	RestaurantId string      `json:"restaurant_id"`
	CustomerId   string      `json:"customer_id"`
	CourierId    string      `json:"courier_id"`
	Address      string      `json:"address"`
	Latitude     float32     `json:"lat"`
	Longitude    float32     `json:"long"`
	Items        []OrderItem `json:"items"`
}

type OrderItem struct {
	ID     string `json:"id"`
	ItemID string `json:"item_id"`
	Qty    uint32 `json:"qty"`
	Price  uint32 `json:"price"`
}

type Courier struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	PhoneNumber string  `json:"phone_number"`
	Photo       string  `json:"photo"`
	CarModel    string  `json:"car_model"`
	CarNumber   string  `json:"car_number"`
	CarColor    string  `json:"car_color"`
	Lat         float32 `json:"lat"`
	Long        float32 `json:"long"`
}
