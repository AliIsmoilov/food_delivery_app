package service

import (
	"context"
	pb "monorepo/src/idl/restaurant_service"
	libsUtils "monorepo/src/libs/utils"
	"monorepo/src/restaurant_service/entity"
	"monorepo/src/restaurant_service/mappers"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (as *RestaurantService) CreateCategory(ctx context.Context, req *pb.Category) (*pb.PrimaryKey, error) {
	as.logger.For(ctx).Info("CreateCategory started",
		zap.Any("Req:", req),
	)
	req.Id = uuid.New().String()
	category, err := mappers.ToCategoryEntity(req)
	if err != nil {
		as.logger.For(ctx).Info("failed in mappers.ToCategoryEntity")
		return &pb.PrimaryKey{}, err
	}

	id, err := as.storage.Category().CreateCategory(ctx, category)
	if err != nil {
		as.logger.For(ctx).Info("failed in storage.CreateCategory")
		return &pb.PrimaryKey{}, err
	}
	as.logger.For(ctx).Info("CreateCategory finished")
	return &pb.PrimaryKey{Id: id.String()}, nil
}

func (as *RestaurantService) ListCategories(ctx context.Context, req *pb.ListCategoriesReq) (*pb.ListCategoriesResp, error) {
	as.logger.For(ctx).Info("ListCategories started",
		zap.Any("Req:", req),
	)

	resp, err := as.storage.Category().ListCategories(ctx, entity.ListCategoriesReq{})
	if err != nil {
		as.logger.For(ctx).Info("failed in storage.ListCategories")
		return &pb.ListCategoriesResp{}, err
	}
	as.logger.For(ctx).Info("ListCategories finished")
	return &pb.ListCategoriesResp{
		Count:      int32(resp.Count),
		Categories: libsUtils.MapSlice(resp.Categories, mappers.ToCategoryProto),
	}, nil
}
