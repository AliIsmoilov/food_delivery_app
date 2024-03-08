package service

import (
	"context"
	pb "monorepo/src/idl/order_service"
	"monorepo/src/order_service/mappers"

	"go.uber.org/zap"
)

func (as *OrderService) CreateCourier(ctx context.Context, req *pb.Courier) (*pb.PrimaryKey, error) {
	as.logger.For(ctx).Info("CreateCourier started",
		zap.Any("Req:", req),
	)

	courer, err := mappers.ToCourierEntity(req)
	if err != nil {
		as.logger.For(ctx).Info("failed in ToCourierEntity")
		return &pb.PrimaryKey{}, err
	}
	id, err := as.storage.Courier().CreateCourier(ctx, courer)
	if err != nil {
		as.logger.For(ctx).Info("failed in storage.CreateCourier")
		return &pb.PrimaryKey{}, err
	}
	as.logger.For(ctx).Info("CreateCourier finished")
	return &pb.PrimaryKey{Id: id.String()}, nil
}
