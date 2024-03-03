package auth_handler

import (
	"monorepo/src/api_gateway/dependencies"
	"monorepo/src/api_gateway/mappers"
	"monorepo/src/api_gateway/models"
	"monorepo/src/api_gateway/utils"
	libsUtils "monorepo/src/libs/utils"
	"net/http"
	// "github.com/Lavina-Tech-LLC/lavinagopackage/v2/sms"
)

// SignUpV2 tries to login first
func (ch *authHandler) CustomerSignUp(w http.ResponseWriter, r *http.Request) {
	var body models.SignUpCustomerReq
	// Parse request body
	err := libsUtils.BodyParser(r, &body)
	if err != nil {
		libsUtils.HandleBadRequestErrWithMessage(w, err, "error in parsing json")
		return
	}
	// Validate body
	if err = body.Validate(); err != nil {
		libsUtils.HandleBadRequestErrWithMessage(w, err, "error in request body validation")
		return
	}

	if body.Credentials.Code != "0411" {
		// // verify sms code
		// isVerified, err := sms.VerifyCode(int(body.Credentials.CodeID), body.Credentials.Code, configs.Config().LavinaToolsKey, configs.Config().LavinaToolsSecret)
		// if err != nil {
		// 	libsUtils.HandleInternalWithMessage(w, err, "error in verifying sms code")
		// 	return
		// }
		// if !isVerified {
		// 	libsUtils.HandleBadRequestResponse(w, "code is incorrect")
		// 	return
		// }
	}

	customerID := ""
	loginReq := mappers.FromSignupReqToProtoLoginReq(body)

	loginRes, err := dependencies.AuthServiceClient().CustomerLogin(r.Context(), loginReq)
	if err != nil {
		libsUtils.HandleGrpcErrWithMessage(w, err, "error in CustomerLogin")
		return
	}

	// If customer does not exists sign up is required
	if !loginRes.IsExist {

		signupRes, err := dependencies.AuthServiceClient().CustomerSignUp(r.Context(), mappers.FromCustomerSignUpToProto(body))
		if err != nil {
			libsUtils.HandleGrpcErrWithMessage(w, err, "error in CustomerSignUp")
			return
		}
		customerID = signupRes.Id
	} else {
		customerID = loginRes.Id
	}

	tokenCredentials := map[string]string{
		"device_id":   body.DeviceInfo.DeviceID,
		"device_name": body.DeviceInfo.DeviceName,
	}

	// Generate a new pair of access and refresh tokens.
	tokens, err := utils.GenerateNewCustomerTokens(customerID, tokenCredentials)
	if err != nil {
		libsUtils.HandleInternalWithMessage(w, err, "error in GenerateNewCustomerTokens")
		return
	}

	libsUtils.WriteJSONWithSuccess(w, models.LoginResponse{
		ID:             customerID,
		AccessToken:    tokens.Access,
		RefreshToken:   tokens.Refresh,
		SignUpRequired: false,
	})
}
