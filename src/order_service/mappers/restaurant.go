package mappers

import (
	pb "monorepo/src/idl/restaurant_service"
	"monorepo/src/restaurant_service/entity"

	"github.com/google/uuid"
)

func ToRestaurantEntity(restaurant *pb.Restaurant) (entity.Restaurant, error) {

	id, err := uuid.Parse(restaurant.Id)
	if err != nil {
		return entity.Restaurant{}, err
	}

	return entity.Restaurant{
		Id:          id,
		Name:        restaurant.Name,
		Adress:      restaurant.Address,
		PhoneNumber: restaurant.PhoneNumber,
		Latitude:    float64(restaurant.Lat),
		Longitude:   float64(restaurant.Lon),
		Photo:       restaurant.Photo,
	}, nil
}

func ToListRestaurantsReqEntity(req *pb.ListRestaurantsReq) entity.ListRestaurantsReq {
	return entity.ListRestaurantsReq{}
}

func ToRestaurantProto(res entity.Restaurant) *pb.Restaurant {
	return &pb.Restaurant{
		Id:          res.Id.String(),
		Name:        res.Name,
		Address:     res.Adress,
		Lat:         float32(res.Latitude),
		Lon:         float32(res.Longitude),
		Photo:       res.Photo,
		PhoneNumber: res.PhoneNumber,
	}
}
