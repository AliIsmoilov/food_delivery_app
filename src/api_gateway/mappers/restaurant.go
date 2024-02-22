package mappers

import (
	"monorepo/src/api_gateway/models"
	pb "monorepo/src/idl/restaurant_service"
)

func ToRestaurantProto(req models.Restaurant) *pb.Restaurant {
	return &pb.Restaurant{
		Id:          req.Id,
		Name:        req.Name,
		Address:     req.Adress,
		Lat:         req.Latitude,
		Lon:         req.Longitude,
		Photo:       req.Photo,
		PhoneNumber: req.PhoneNumber,
	}
}

func ToCategoryProto(req models.Category) *pb.Category {
	return &pb.Category{
		Id:   req.Id,
		Name: req.Name,
	}
}

func ToItemProto(req models.Item) *pb.Item {
	return &pb.Item{
		Id:           req.Id,
		Name:         req.Name,
		CategoryId:   req.CategoryId,
		RestaurantId: req.RestaurantId,
		Description:  req.Description,
		Photo:        req.Photo,
		Price:        req.Price,
	}
}
