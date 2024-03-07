package order

import (
	"monorepo/src/api_gateway/dependencies"
	"monorepo/src/api_gateway/mappers"
	"monorepo/src/api_gateway/models"
	"monorepo/src/idl/order_service"
	"monorepo/src/libs/log"
	libsUtils "monorepo/src/libs/utils"
	"net/http"
)

// RestaurantHandler ...
type OrderHandler struct {
	logger      log.Factory
	orderClient order_service.OrderServiceClient
}

// New creates handler
func New(logger log.Factory) OrderHandler {
	return OrderHandler{
		logger:      logger,
		orderClient: dependencies.OrderServiceClient(),
	}
}

// CreateRestaurant ...
func (rh *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var body models.Order
	err := libsUtils.BodyParser(r, &body)
	if err != nil {
		libsUtils.HandleBadRequestErrWithMessage(w, err, "error in parsing json: ")
		return
	}

	res, err := dependencies.OrderServiceClient().CreateOrder(r.Context(), mappers.ToOrderProto(body))
	if err != nil {
		libsUtils.HandleGrpcErrWithMessage(w, err, "Error at CreateOrder")
		return
	}

	libsUtils.WriteJSONWithSuccess(w, res)
}

func (rh *OrderHandler) OrderList(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	// limit, _ := strconv.Atoi(qryParams.Get("limit"))
	// page, _ := strconv.Atoi(qryParams.Get("page"))

	resp, err := dependencies.OrderServiceClient().OrderList(r.Context(), &order_service.OrderListReq{
		Status: q.Get("status"),
	})
	if err != nil {
		libsUtils.HandleGrpcErrWithMessage(w, err, "Error at OrderList")
		return
	}

	libsUtils.WriteJSONWithSuccess(w, resp)
}
