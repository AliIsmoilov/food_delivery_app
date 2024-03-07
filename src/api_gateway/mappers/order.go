package mappers

import (
	"monorepo/src/api_gateway/models"
	pb "monorepo/src/idl/order_service"
	libsUtils "monorepo/src/libs/utils"
)

func ToOrderProto(req models.Order) *pb.Order {
	return &pb.Order{
		Id:           req.Id,
		Code:         req.Code,
		Note:         req.Note,
		TotalPrice:   req.TotalPrice,
		Status:       req.Status,
		RestaurantId: req.RestaurantId,
		CustomerId:   req.CustomerId,
		CourierId:    req.CourierId,
		Adress:       req.Address,
		Lat:          req.Latitude,
		Long:         req.Longitude,
		Items:        libsUtils.MapSlice(req.Items, ToOrderItem),
	}
}

func ToOrderItem(item models.OrderItem) *pb.Item {
	return &pb.Item{
		ItemId: item.ItemID,
		Qty:    item.Qty,
	}
}
