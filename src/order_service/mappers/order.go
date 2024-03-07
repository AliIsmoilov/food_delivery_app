package mappers

import (
	pb "monorepo/src/idl/order_service"
	libsUtils "monorepo/src/libs/utils"
	"monorepo/src/order_service/entity"

	"github.com/google/uuid"
)

func ToOrdertEntity(order *pb.Order) entity.Order {

	id := uuid.New()
	return entity.Order{
		Id:   id,
		Code: order.Code,
		Note: order.Note,
		// TotalPrice:   order.TotalPrice,
		TotalPrice:   100,
		Status:       order.Status,
		RestaurantId: order.RestaurantId,
		CustomerId:   order.CustomerId,
		CourierId:    order.CourierId,
		Address:      order.Adress,
		Lat:          float64(order.Lat),
		Long:         float64(order.Long),
		Items:        ToItemsEntity(order.Items, id.String()),
	}
}

func ToItemsEntity(items []*pb.Item, orderId string) []entity.Item {
	resp := []entity.Item{}
	for i, item := range items {
		resp = append(resp, ToItemEntity(item))
		resp[i].OrderID = orderId
	}
	return resp
}

func ToItemEntity(item *pb.Item) entity.Item {
	return entity.Item{
		ID:     uuid.New().String(),
		ItemID: item.ItemId,
		Qty:    item.Qty,
		Price:  10,
	}
}

// func ToListRestaurantsReqEntity(req *pb.ListRestaurantsReq) entity.ListRestaurantsReq {
// 	return entity.ListRestaurantsReq{}
// }

func ToOrderProto(order entity.Order) *pb.Order {
	return &pb.Order{
		Id:           order.Id.String(),
		Code:         order.Code,
		Note:         order.Note,
		TotalPrice:   order.TotalPrice,
		Status:       order.Status,
		RestaurantId: order.RestaurantId,
		CustomerId:   order.CustomerId,
		CourierId:    order.CourierId,
		Adress:       order.Address,
		Lat:          float32(order.Lat),
		Long:         float32(order.Long),
		Items:        libsUtils.MapSlice(order.Items, ToOrderItemProto),
	}
}

func ToOrderItemProto(item entity.Item) *pb.Item {
	return &pb.Item{
		Id:      item.ID,
		ItemId:  item.ItemID,
		OrderId: item.OrderID,
		Qty:     item.Qty,
		Price:   item.Price,
	}
}

func ToOrderListReq(req *pb.OrderListReq) entity.OrderListReq {
	return entity.OrderListReq{
		Status: req.Status,
	}
}
