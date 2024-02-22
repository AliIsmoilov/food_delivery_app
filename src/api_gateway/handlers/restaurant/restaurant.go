package restaurant

import (
	"monorepo/src/api_gateway/dependencies"
	"monorepo/src/api_gateway/mappers"
	"monorepo/src/api_gateway/models"
	"monorepo/src/idl/restaurant_service"
	"monorepo/src/libs/log"
	libsUtils "monorepo/src/libs/utils"
	"net/http"
)

// RestaurantHandler ...
type RestaurantHandler struct {
	logger           log.Factory
	restaurantClient restaurant_service.RestaurantServiceClient
}

// New creates handler
func New(logger log.Factory) RestaurantHandler {
	return RestaurantHandler{
		logger:           logger,
		restaurantClient: dependencies.RestauranticeClient(),
	}
}

// CreateRestaurant ...
func (rh *RestaurantHandler) CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	var body models.Restaurant
	err := libsUtils.BodyParser(r, &body)
	if err != nil {
		libsUtils.HandleBadRequestErrWithMessage(w, err, "error in parsing json: ")
		return
	}

	res, err := rh.restaurantClient.CreateRestaurant(r.Context(), mappers.ToRestaurantProto(body))
	if err != nil {
		libsUtils.HandleGrpcErrWithMessage(w, err, "Error at branch created")
		return
	}

	libsUtils.WriteJSONWithSuccess(w, res)
}

func (rh *RestaurantHandler) ListRestaurants(w http.ResponseWriter, r *http.Request) {

	// qryParams := r.URL.Query()

	// limit, _ := strconv.Atoi(qryParams.Get("limit"))
	// page, _ := strconv.Atoi(qryParams.Get("page"))

	resp, err := rh.restaurantClient.ListRestaurants(r.Context(), &restaurant_service.ListRestaurantsReq{})
	if err != nil {
		libsUtils.HandleGrpcErrWithMessage(w, err, "Error at ListWarehouseCategories")
		return
	}

	libsUtils.WriteJSONWithSuccess(w, resp)
}
