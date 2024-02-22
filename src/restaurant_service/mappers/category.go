package mappers

import (
	pb "monorepo/src/idl/restaurant_service"
	"monorepo/src/restaurant_service/entity"

	"github.com/google/uuid"
)

func ToCategoryEntity(req *pb.Category) (entity.Category, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return entity.Category{}, err
	}

	return entity.Category{
		Id:   id,
		Name: req.Name,
	}, nil
}

func ToCategoryProto(req entity.Category) *pb.Category {
	return &pb.Category{
		Id:   req.Id.String(),
		Name: req.Name,
	}
}
