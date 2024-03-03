package mappers

import (
	"monorepo/src/auth_service/pkg/entity"
	pb "monorepo/src/idl/auth_service"

	"github.com/google/uuid"
)

func FromProtoToLoginRequest(req *pb.CustLoginRequest) (entity.LoginRequest, error) {

	return entity.LoginRequest{
		PhoneNumber: req.GetPhoneNumber(),
		DeviceInfo: entity.DeviceInfo{
			DeviceId:       req.Device.DeviceId,
			DeviceName:     req.Device.DeviceName,
			NotificationId: req.Device.DeviceId,
		},
	}, nil
}

// ToProtoCreateResponse ...
func ToProtoCreateResponse(req entity.CreateResponse) *pb.CreateCustomerResponse {
	return &pb.CreateCustomerResponse{
		Id:      req.ID,
		Name:    req.Name,
		Photo:   req.Photo,
		Email:   req.Email,
		IsExist: req.IsExist,
	}
}

// MapProtoToCustomerReq ...
func MapProtoToCustomerReq(req *pb.CreateCustomerRequest) entity.Customer {
	customerID := uuid.New()
	return entity.Customer{
		ID:          customerID,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Email:       req.Email,
		Photo:       req.Photo,
		DeviceInfo: entity.DeviceInfo{
			Id:             uuid.New(),
			CustomerId:     customerID,
			DeviceId:       req.DeviceInfo.DeviceId,
			DeviceName:     req.DeviceInfo.DeviceName,
			NotificationId: req.DeviceInfo.NotificationId,
		},
	}
}

// FromCustomerEntityToCreateResponse ...
func FromCustomerEntityToCreateResponse(req entity.Customer) *pb.CreateCustomerResponse {
	return &pb.CreateCustomerResponse{
		Id:    req.ID.String(),
		Name:  req.Name,
		Photo: req.Photo,
		Email: req.Email,
	}
}
