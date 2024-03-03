package models

import (
	"fmt"
	"regexp"
	"strings"

	libCons "monorepo/src/libs/constants"
	libsUtils "monorepo/src/libs/utils"

	validation "github.com/go-ozzo/ozzo-validation/v3"
)

// StuffLoginModel ...
type StuffLoginModel struct {
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

// Validate StuffLoginModel
func (lm *StuffLoginModel) Validate() error {
	return validation.ValidateStruct(
		lm,
		validation.Field(&lm.Password, validation.Required, validation.Length(8, 30), validation.Match(regexp.MustCompile("[a-z]|[A-Z][0-9]"))),
		validation.Field(&lm.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile("[+]{1}[1-9]{1}[0-9]{1,13}$"))),
	)
}

type LoginResponse struct {
	ID             string `json:"id"`
	AccessToken    string `json:"access_token"`
	RefreshToken   string `json:"refresh_token"`
	SignUpRequired bool   `json:"sign_up_required"`
}

type ReqResetPassword struct {
	NewPassword string `json:"new_password"`
}

// UserInfo ...
type UserInfo struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Email       string `json:"email,omitempty"`
	Photo       string `json:"photo,omitempty"`
	Language    string `json:"language,omitempty"`
	Birthdate   int64  `json:"birthdate,omitempty"`
	Sex         string `json:"sex,omitempty"`
	SocialID    string `json:"social_id,omitempty"`
}

// DeviceInfo ...
type DeviceInfo struct {
	DeviceID       string `json:"device_id"`
	DeviceName     string `json:"device_name"`
	NotificationID string `json:"notification_id"`
}

// Credentials ...
type Credentials struct {
	CodeID uint64 `json:"code_id"`
	Code   string `json:"code"`
}

// SignUpCustomerReq ...
type SignUpCustomerReq struct {
	Credentials Credentials `json:"credentials"`
	UserInfo    UserInfo    `json:"user"`
	DeviceInfo  DeviceInfo  `json:"device"`
	From        string      `json:"from"` //can be "mobile, telegram, web"
}

// Validate ...
func (r *SignUpCustomerReq) Validate() error {

	if ok := libsUtils.IsPhoneValid(r.UserInfo.PhoneNumber); !ok {
		return fmt.Errorf("%w: invalid phone number", libCons.ErrInvalidRequestBody)
	}

	if r.UserInfo.Email != "" {
		if ok := libsUtils.IsEmailValid(r.UserInfo.Email); !ok {
			return fmt.Errorf("%w: invalid email", libCons.ErrInvalidRequestBody)
		}
	}

	// if r.Credentials.Code == "" {
	// 	return fmt.Errorf("%w: invalid creadential values", libCons.ErrInvalidRequestBody)
	// }

	if len(strings.TrimSpace(r.UserInfo.Name)) < 3 && len(strings.TrimSpace(r.UserInfo.Name)) > 255 {
		return libCons.ErrIncorrectNameValue
	}

	return nil
}
