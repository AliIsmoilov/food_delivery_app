package service

import (
	"fmt"
	"monorepo/src/libs/log"
	"monorepo/src/order_service/storage"

	"github.com/opentracing/opentracing-go"
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

// func (as *RestaurantService) CreateRestaurant(ctx context.Context, req *pb.Restaurant) (*pb.PrimaryKey, error) {
// 	as.logger.For(ctx).Info("CreateRestaurant started",
// 		zap.Any("Req:", req),
// 	)
// 	req.Id = uuid.New().String()
// 	restaurant, err := mappers.ToRestaurantEntity(req)
// 	if err != nil {
// 		as.logger.For(ctx).Info("failed in mappers.ToRestaurantEntity")
// 		return &pb.PrimaryKey{}, err
// 	}

// 	id, err := as.storage.Restaurant().CreateRestaurant(ctx, restaurant)
// 	if err != nil {
// 		as.logger.For(ctx).Info("failed in storage.CreateRestaurant")
// 		return &pb.PrimaryKey{}, err
// 	}
// 	as.logger.For(ctx).Info("CreateRestaurant finished")
// 	return &pb.PrimaryKey{Id: id.String()}, nil
// }

// func (as *RestaurantService) ListRestaurants(ctx context.Context, req *pb.ListRestaurantsReq) (*pb.ListRestaurantsResp, error) {
// 	as.logger.For(ctx).Info("ListRestaurants started",
// 		zap.Any("Req:", req),
// 	)

// 	resp, err := as.storage.Restaurant().ListRestaurants(ctx, mappers.ToListRestaurantsReqEntity(req))
// 	if err != nil {
// 		as.logger.For(ctx).Info("failed in storage.ListRestaurant")
// 		return &pb.ListRestaurantsResp{}, err
// 	}
// 	as.logger.For(ctx).Info("ListRestaurants finished")
// 	return &pb.ListRestaurantsResp{
// 		Count:       int32(resp.Count),
// 		Restaurants: libsUtils.MapSlice(resp.Restaurants, mappers.ToRestaurantProto),
// 	}, nil
// }
