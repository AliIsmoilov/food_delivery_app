package order

import (
	"fmt"
	"monorepo/src/api_gateway/dependencies"
	"monorepo/src/api_gateway/mappers"
	"monorepo/src/api_gateway/models"
	"monorepo/src/idl/order_service"
	libsUtils "monorepo/src/libs/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (rh *OrderHandler) CreateCourier(w http.ResponseWriter, r *http.Request) {
	var body models.Courier
	err := libsUtils.BodyParser(r, &body)
	if err != nil {
		libsUtils.HandleBadRequestErrWithMessage(w, err, "error in parsing json: ")
		return
	}

	res, err := dependencies.OrderServiceClient().CreateCourier(r.Context(), mappers.ToCourierProto(body))
	if err != nil {
		libsUtils.HandleGrpcErrWithMessage(w, err, "Error at CreateCourier")
		return
	}

	libsUtils.WriteJSONWithSuccess(w, res)
}

func (rh *OrderHandler) CourierList(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	limit, _ := strconv.Atoi(q.Get("limit"))
	page, _ := strconv.Atoi(q.Get("page"))

	resp, err := dependencies.OrderServiceClient().CourierList(r.Context(), &order_service.CourierListReq{
		Limit:  uint32(limit),
		Page:   uint32(page),
		Search: q.Get("search"),
	})
	if err != nil {
		libsUtils.HandleGrpcErrWithMessage(w, err, "Error at CourierList")
		return
	}

	libsUtils.WriteJSONWithSuccess(w, resp)
}

func (wh *OrderHandler) UpdateCourier(w http.ResponseWriter, r *http.Request) {
	var body models.Courier
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		libsUtils.HandleBadRequestErrWithMessage(w, fmt.Errorf("id is missing"), " ")
		return
	}

	if !libsUtils.IsUUID(id) {
		libsUtils.HandleBadRequestErrWithMessage(w, fmt.Errorf("id is missing"), " ")
		return
	}
	if err := libsUtils.BodyParser(r, &body); err != nil {
		libsUtils.HandleBadRequestErrWithMessage(w, err, "error in parsing json")
		return
	}
	body.ID = id

	_, err := dependencies.OrderServiceClient().UpdateCourier(r.Context(), mappers.ToCourierProto(body))
	if err != nil {
		libsUtils.HandleGrpcErrWithMessage(w, err, "Error at UpdateCourier")
		return
	}

	libsUtils.WriteJSONWithSuccess(w, http.StatusOK)
}

func (wh *OrderHandler) DeleteCourier(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		libsUtils.HandleBadRequestErrWithMessage(w, fmt.Errorf("id is missing"), " ")
		return
	}
	if !libsUtils.IsUUID(id) {
		libsUtils.HandleBadRequestResponse(w, "id is not valid uuid")
		return
	}

	_, err := dependencies.OrderServiceClient().DeleteCourier(r.Context(), &order_service.PrimaryKey{
		Id: id,
	})
	if err != nil {
		libsUtils.HandleGrpcErrWithMessage(w, err, "Error in DeleteCourier")
		return
	}

	libsUtils.WriteJSONWithSuccess(w, http.StatusOK)
}

func (wh *OrderHandler) GetCourier(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		libsUtils.HandleBadRequestErrWithMessage(w, fmt.Errorf("id is missing"), " ")
		return
	}
	if !libsUtils.IsUUID(id) {
		libsUtils.HandleBadRequestResponse(w, "seller_id is not valid uuid")
		return
	}

	resp, err := dependencies.OrderServiceClient().GetCourier(r.Context(), &order_service.PrimaryKey{
		Id: id,
	})
	if err != nil {
		libsUtils.HandleGrpcErrWithMessage(w, err, "Error in GetCourier")
		return
	}

	libsUtils.WriteJSONWithSuccess(w, resp)
}
