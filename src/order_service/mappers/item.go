package mappers

import (
	pb "monorepo/src/idl/restaurant_service"
	libsUtils "monorepo/src/libs/utils"
	"monorepo/src/restaurant_service/entity"

	"github.com/google/uuid"
)

func ToItemEntity(item *pb.Item) (entity.Item, error) {

	id, err := uuid.Parse(item.Id)
	if err != nil {
		return entity.Item{}, err
	}

	return entity.Item{
		Id:           id,
		Name:         item.Name,
		CategoryId:   item.CategoryId,
		RestaurantId: item.RestaurantId,
		Description:  item.Description,
		Photo:        item.Photo,
		Price:        item.Price,
	}, nil
}

func ToRestaurantItemsGroupedByCategoryReq(req *pb.RestaurantItemsGroupedByCategoryReq) (entity.RestaurantItemsGroupedByCategoryReq, error) {
	id, err := uuid.Parse(req.RestaurantId)
	if err != nil {
		return entity.RestaurantItemsGroupedByCategoryReq{}, err
	}
	return entity.RestaurantItemsGroupedByCategoryReq{
		RestaurantId: id,
	}, nil
}

func ToItemProto(item entity.Item) *pb.Item {
	return &pb.Item{
		Id:           item.Id.String(),
		Name:         item.Name,
		CategoryId:   item.CategoryId,
		RestaurantId: item.RestaurantId,
		Description:  item.Description,
		Photo:        item.Photo,
		Price:        item.Price,
	}
}

func ToRestaurantItemsGroupedByCategoryModel(req entity.RestaurantItemsGroupedByCategoryModel) *pb.RestaurantItemsModel {
	return &pb.RestaurantItemsModel{
		Category: ToCategoryProto(req.Category),
		Items:    libsUtils.MapSlice(req.Items, ToItemProto),
	}
}
