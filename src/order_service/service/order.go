package service

import (
	"context"
	"fmt"
	pb "monorepo/src/idl/order_service"
	"monorepo/src/libs/log"
	libsUtils "monorepo/src/libs/utils"
	"monorepo/src/order_service/mappers"
	"monorepo/src/order_service/storage"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type OrderService struct {
	storage storage.IStorage
	logger  log.Factory
	tracer  opentracing.Tracer
}

// NewUserService ...
func New(logger log.Factory, tracer opentracing.Tracer, storage storage.IStorage) *OrderService {
	fmt.Println("Service new")
	orderServer := &OrderService{
		storage: storage,
		logger:  logger,
		tracer:  tracer,
	}
	return orderServer
}

func (as *OrderService) CreateOrder(ctx context.Context, req *pb.Order) (*pb.PrimaryKey, error) {
	as.logger.For(ctx).Info("CreateOrder started",
		zap.Any("Req:", req),
	)

	id, err := as.storage.Order().CreateOrder(ctx, mappers.ToOrdertEntity(req))
	if err != nil {
		as.logger.For(ctx).Info("failed in storage.CreateOrder")
		return &pb.PrimaryKey{}, err
	}
	as.logger.For(ctx).Info("CreateOrder finished")
	return &pb.PrimaryKey{Id: id.String()}, nil
}

func (as *OrderService) OrderList(ctx context.Context, req *pb.OrderListReq) (*pb.OrderListResp, error) {
	as.logger.For(ctx).Info("OrderList started",
		zap.Any("Req:", req),
	)

	resp, err := as.storage.Order().OrderList(ctx, mappers.ToOrderListReq(req))
	if err != nil {
		as.logger.For(ctx).Info("failed in storage.OrderList")
		return &pb.OrderListResp{}, err
	}
	as.logger.For(ctx).Info("OrderList finished")
	return &pb.OrderListResp{
		Count:  uint32(resp.Count),
		Orders: libsUtils.MapSlice(resp.Orders, mappers.ToOrderProto),
	}, nil
}
