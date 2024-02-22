package dependencies

import (
	"fmt"
	"monorepo/src/api_gateway/configs"
	"monorepo/src/idl/auth_service"
	"monorepo/src/idl/restaurant_service"
	"sync"

	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	authServiceClient           auth_service.AuthServiceClient
	restaurantServiceClient     restaurant_service.RestaurantServiceClient
	onceAuthServiceClient       sync.Once
	onceRestaurantServiceClient sync.Once
)

// AuthServiceClient loads AuthServiceClient using atomic pattern
func AuthServiceClient() auth_service.AuthServiceClient {
	onceAuthServiceClient.Do(func() {
		authServiceClient = loadAuthServiceClient()
	})
	return authServiceClient
}

// RestaurantClient loads RestaurantClient using atomic pattern
func RestauranticeClient() restaurant_service.RestaurantServiceClient {
	onceRestaurantServiceClient.Do(func() {
		restaurantServiceClient = loadRestaurantServiceClient()
	})
	return restaurantServiceClient
}

func loadAuthServiceClient() auth_service.AuthServiceClient {
	tracer := opentracing.GlobalTracer()
	conf := configs.Config()
	connAuth, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.AuthServiceHost, conf.AuthServicePort),
		grpc.WithTransportCredentials(
			insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(tracer)),
		grpc.WithStreamInterceptor(
			otgrpc.OpenTracingStreamClientInterceptor(tracer),
		),
	)
	if err != nil {
		panic(fmt.Errorf("auth service dial host: %s port:%d err: %s",
			conf.AuthServiceHost, conf.AuthServicePort, err))
	}

	return auth_service.NewAuthServiceClient(connAuth)
}

func loadRestaurantServiceClient() restaurant_service.RestaurantServiceClient {
	tracer := opentracing.GlobalTracer()
	conf := configs.Config()
	connRestaurant, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.RestaurantServiceHost, conf.RestaurantServicePort),
		grpc.WithTransportCredentials(
			insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(tracer)),
		grpc.WithStreamInterceptor(
			otgrpc.OpenTracingStreamClientInterceptor(tracer),
		),
	)
	if err != nil {
		panic(fmt.Errorf("auth service dial host: %s port:%d err: %s",
			conf.AuthServiceHost, conf.AuthServicePort, err))
	}

	return restaurant_service.NewRestaurantServiceClient(connRestaurant)
}
