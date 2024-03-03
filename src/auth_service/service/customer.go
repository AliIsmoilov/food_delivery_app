package service

import (
	"context"

	"monorepo/src/auth_service/pkg/mappers"
	pb "monorepo/src/idl/auth_service"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Check Customer is exist by phoneNumber or Email and return id, ... & true flag if exists
func (s *AuthService) CustomerLogin(ctx context.Context, req *pb.CustLoginRequest) (*pb.CreateCustomerResponse, error) {
	s.logger.For(ctx).Info("Customer login started")

	request, err := mappers.FromProtoToLoginRequest(req)
	if err != nil {
		s.logger.For(ctx).Error("error in mapping proto request into entity", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := s.storage.Customer().CustomerLogin(ctx, request)
	if err != nil {
		s.logger.For(ctx).Error("failed in CustomerLogin", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	s.logger.For(ctx).Info("Customer login finished")

	return mappers.ToProtoCreateResponse(resp), nil
}

// Signing up customer
func (s *AuthService) CustomerSignUp(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	s.logger.For(ctx).Info("CustomerSignUp started", zap.String("PhoneNumber", req.PhoneNumber))

	res, err := s.storage.Customer().SignUp(ctx, mappers.MapProtoToCustomerReq(req))
	if err != nil {
		s.logger.For(ctx).Error("Failed in SignUp", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	s.logger.For(ctx).Info("CustomerSignUp finished")
	return mappers.FromCustomerEntityToCreateResponse(res), nil
}
