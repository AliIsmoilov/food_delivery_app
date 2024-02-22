package restaurant

import (
	"monorepo/src/api_gateway/mappers"
	"monorepo/src/api_gateway/models"
	"monorepo/src/idl/restaurant_service"
	libsUtils "monorepo/src/libs/utils"
	"net/http"
)

// CreateRestaurant ...
func (rh *RestaurantHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var body models.Item
	err := libsUtils.BodyParser(r, &body)
	if err != nil {
		libsUtils.HandleBadRequestErrWithMessage(w, err, "error in parsing json: ")
		return
	}

	res, err := rh.restaurantClient.CreateItem(r.Context(), mappers.ToItemProto(body))
	if err != nil {
		libsUtils.HandleGrpcErrWithMessage(w, err, "Error at CreateItem")
		return
	}

	libsUtils.WriteJSONWithSuccess(w, res)
}

func (rh *RestaurantHandler) RestaurantItemsGroupedByCategory(w http.ResponseWriter, r *http.Request) {

	qryParams := r.URL.Query()
	restaurantId := qryParams.Get("restaurant_id")

	resp, err := rh.restaurantClient.RestaurantItemsGroupedByCategory(r.Context(), &restaurant_service.RestaurantItemsGroupedByCategoryReq{
		RestaurantId: restaurantId,
	})
	if err != nil {
		libsUtils.HandleGrpcErrWithMessage(w, err, "Error at RestaurantItemsGroupedByCategory")
		return
	}

	libsUtils.WriteJSONWithSuccess(w, resp)
}
