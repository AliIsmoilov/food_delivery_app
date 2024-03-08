package service

import (
	"context"
	pb "monorepo/src/idl/order_service"
	libsUtils "monorepo/src/libs/utils"
	"monorepo/src/order_service/mappers"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (as *OrderService) CreateCourier(ctx context.Context, req *pb.Courier) (*pb.PrimaryKey, error) {
	as.logger.For(ctx).Info("CreateCourier started",
		zap.Any("Req:", req),
	)

	courer, err := mappers.ToCourierEntity(req)
	courer.ID = uuid.New()
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

func (as *OrderService) CourierList(ctx context.Context, req *pb.CourierListReq) (*pb.CourierListResp, error) {
	as.logger.For(ctx).Info("CourierList started",
		zap.Any("Req:", req),
	)

	resp, err := as.storage.Courier().CourierList(ctx, mappers.ToCourierListReq(req))
	if err != nil {
		as.logger.For(ctx).Info("failed in storage.CourierList")
		return &pb.CourierListResp{}, err
	}
	as.logger.For(ctx).Info("CourierList finished")
	return &pb.CourierListResp{
		Count:    int32(resp.Count),
		Couriers: libsUtils.MapSlice(resp.Couriers, mappers.ToCourierProto),
	}, nil
}

func (as *OrderService) UpdateCourier(ctx context.Context, req *pb.Courier) (*pb.Empty, error) {
	as.logger.For(ctx).Info("UpdateCourier started",
		zap.Any("Req:", req),
	)

	courer, err := mappers.ToCourierEntity(req)
	if err != nil {
		as.logger.For(ctx).Info("failed in ToCourierEntity")
		return &pb.Empty{}, err
	}
	err = as.storage.Courier().UpdateCourier(ctx, courer)
	if err != nil {
		as.logger.For(ctx).Info("failed in storage.UpdateCourier")
		return &pb.Empty{}, err
	}
	as.logger.For(ctx).Info("UpdateCourier finished")
	return &pb.Empty{}, nil
}

func (as *OrderService) DeleteCourier(ctx context.Context, req *pb.PrimaryKey) (*pb.Empty, error) {
	as.logger.For(ctx).Info("DeleteCourier started",
		zap.Any("Req:", req),
	)

	err := as.storage.Courier().DeleteCourier(ctx, req.Id)
	if err != nil {
		as.logger.For(ctx).Info("failed in storage.DeleteCourier")
		return &pb.Empty{}, err
	}
	as.logger.For(ctx).Info("DeleteCourier finished")
	return &pb.Empty{}, nil
}

func (as *OrderService) GetCourier(ctx context.Context, req *pb.PrimaryKey) (*pb.Courier, error) {
	as.logger.For(ctx).Info("GetCourier started",
		zap.Any("Req:", req),
	)

	courier, err := as.storage.Courier().GetCourier(ctx, req.Id)
	if err != nil {
		as.logger.For(ctx).Info("failed in storage.GetCourier")
		return &pb.Courier{}, err
	}
	as.logger.For(ctx).Info("GetCourier finished")
	return mappers.ToCourierProto(courier), nil
}
