package mappers

import (
	pb "monorepo/src/idl/order_service"
	_ "monorepo/src/libs/utils"
	"monorepo/src/order_service/entity"

	"github.com/google/uuid"
)

func ToCourierEntity(req *pb.Courier) (entity.Courier, error) {
	id := uuid.UUID{}
	var err error
	if req.Id != "" {
		id, err = uuid.Parse(req.Id)
		if err != nil {
			return entity.Courier{}, err
		}
	}
	return entity.Courier{
		ID:          id,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Photo:       req.Photo,
		CarModel:    req.CarModel,
		CarNumber:   req.CarNumber,
		CarColor:    req.CarColor,
	}, nil
}

func ToCourierProto(req entity.Courier) *pb.Courier {
	return &pb.Courier{
		Id:          req.ID.String(),
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Photo:       req.Photo,
		CarModel:    req.CarModel,
		CarNumber:   req.CarNumber,
		CarColor:    req.CarColor,
		IsAvailable: req.IsAvailable,
		Lat:         req.Latitude,
		Long:        req.Longitude,
	}
}

func ToCourierListReq(req *pb.CourierListReq) entity.CourierListReq {
	return entity.CourierListReq{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
	}
}
