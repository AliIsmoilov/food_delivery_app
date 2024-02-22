package server

import (
	"fmt"
	"monorepo/src/restaurant_service/configs"
	"monorepo/src/restaurant_service/service"
	"net"

	pb "monorepo/src/idl/restaurant_service"

	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

func Start(config *configs.Configuration, tracer opentracing.Tracer, restaurantServer *service.RestaurantService) {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(
		otgrpc.OpenTracingServerInterceptor(tracer)),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(tracer)))

	pb.RegisterRestaurantServiceServer(grpcServer, restaurantServer)

	fmt.Println("Listen: ")

	//listenting tcp rpcport
	lis, err := net.Listen("tcp", config.RPCPort)
	if err != nil {
		fmt.Println("listening tcp error: ", err)
	}

	fmt.Println("crm server running on port : ", config.RPCPort)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("failed to serve: ", err)
	}
}
