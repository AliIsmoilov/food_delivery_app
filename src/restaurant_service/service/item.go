package service

import (
	"context"
	pb "monorepo/src/idl/restaurant_service"
	libsUtils "monorepo/src/libs/utils"
	"monorepo/src/restaurant_service/mappers"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (as *RestaurantService) CreateItem(ctx context.Context, req *pb.Item) (*pb.PrimaryKey, error) {
	as.logger.For(ctx).Info("CreateItem started",
		zap.Any("Req:", req),
	)
	req.Id = uuid.New().String()
	item, err := mappers.ToItemEntity(req)
	if err != nil {
		as.logger.For(ctx).Info("failed in mappers.ToCategoryEntity")
		return &pb.PrimaryKey{}, err
	}

	id, err := as.storage.Item().CreateItem(ctx, item)
	if err != nil {
		as.logger.For(ctx).Info("failed in storage.CreateItem")
		return &pb.PrimaryKey{}, err
	}
	as.logger.For(ctx).Info("CreateItem finished")
	return &pb.PrimaryKey{Id: id.String()}, nil
}

func (as *RestaurantService) RestaurantItemsGroupedByCategory(ctx context.Context, req *pb.RestaurantItemsGroupedByCategoryReq) (*pb.RestaurantItemsGroupedByCategoryResp, error) {
	as.logger.For(ctx).Info("RestaurantItemsGroupedByCategory started",
		zap.Any("Req:", req),
	)

	reqEntity, err := mappers.ToRestaurantItemsGroupedByCategoryReq(req)
	if err != nil {
		as.logger.For(ctx).Info("failed in mappers.ToRestaurantItemsGroupedByCategoryReq")
		return &pb.RestaurantItemsGroupedByCategoryResp{}, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err := as.storage.Item().RestaurantItemsGroupedByCategory(ctx, reqEntity)
	if err != nil {
		as.logger.For(ctx).Info("failed in storage.RestaurantItemsGroupedByCategory")
		return &pb.RestaurantItemsGroupedByCategoryResp{}, status.Error(codes.InvalidArgument, err.Error())
	}
	as.logger.For(ctx).Info("RestaurantItemsGroupedByCategory finished")
	return &pb.RestaurantItemsGroupedByCategoryResp{
		Items: libsUtils.MapSlice(resp.Items, mappers.ToRestaurantItemsGroupedByCategoryModel),
	}, nil
}
