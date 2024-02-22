package restaurant

import (
	"monorepo/src/api_gateway/mappers"
	"monorepo/src/api_gateway/models"
	"monorepo/src/idl/restaurant_service"
	libsUtils "monorepo/src/libs/utils"
	"net/http"
)

// CreateRestaurant ...
func (rh *RestaurantHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var body models.Category
	err := libsUtils.BodyParser(r, &body)
	if err != nil {
		libsUtils.HandleBadRequestErrWithMessage(w, err, "error in parsing json: ")
		return
	}

	res, err := rh.restaurantClient.CreateCategory(r.Context(), mappers.ToCategoryProto(body))
	if err != nil {
		libsUtils.HandleGrpcErrWithMessage(w, err, "Error at CreateCategory")
		return
	}

	libsUtils.WriteJSONWithSuccess(w, res)
}

func (rh *RestaurantHandler) ListCategories(w http.ResponseWriter, r *http.Request) {

	// qryParams := r.URL.Query()

	// limit, _ := strconv.Atoi(qryParams.Get("limit"))
	// page, _ := strconv.Atoi(qryParams.Get("page"))

	resp, err := rh.restaurantClient.ListCategories(r.Context(), &restaurant_service.ListCategoriesReq{})
	if err != nil {
		libsUtils.HandleGrpcErrWithMessage(w, err, "Error at ListCategories")
		return
	}

	libsUtils.WriteJSONWithSuccess(w, resp)
}
