package mappers

import (
	"monorepo/src/api_gateway/models"
	pb "monorepo/src/idl/auth_service"
)

// FromSignupReqToProtoLoginReq ...
func FromSignupReqToProtoLoginReq(body models.SignUpCustomerReq) *pb.CustLoginRequest {
	proto := &pb.CustLoginRequest{
		Device: &pb.DeviceInfo{
			DeviceId:       body.DeviceInfo.DeviceID,
			DeviceName:     body.DeviceInfo.DeviceName,
			NotificationId: body.DeviceInfo.NotificationID,
		},
	}
	if body.UserInfo.PhoneNumber != "" {
		proto.By = &pb.CustLoginRequest_PhoneNumber{
			PhoneNumber: body.UserInfo.PhoneNumber,
		}
		return proto
	} else if body.UserInfo.Email != "" {
		proto.By = &pb.CustLoginRequest_Email{
			Email: body.UserInfo.Email,
		}
		return proto
	}

	return proto
}

func FromCustomerSignUpToProto(body models.SignUpCustomerReq) *pb.CreateCustomerRequest {
	return &pb.CreateCustomerRequest{
		PhoneNumber: body.UserInfo.PhoneNumber,
		Name:        body.UserInfo.Name,
		Email:       body.UserInfo.Email,
		Photo:       body.UserInfo.Photo,
		Language:    body.UserInfo.Language,
		Birthdate:   body.UserInfo.Birthdate,
		Sex:         body.UserInfo.Sex,
		SocialId:    body.UserInfo.SocialID,
		DeviceInfo: &pb.DeviceInfo{
			DeviceId:       body.DeviceInfo.DeviceID,
			DeviceName:     body.DeviceInfo.DeviceName,
			NotificationId: body.DeviceInfo.NotificationID,
		},
	}
}
