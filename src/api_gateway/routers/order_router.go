package routers

import (
	"monorepo/src/api_gateway/handlers"
	"monorepo/src/api_gateway/pkg/tracing"
	"net/http"
)

func RegisterOrderRoutes(r *tracing.TracedServeMux, h *handlers.Handlers) {
	r.Handle("POST", "/order/", http.HandlerFunc(h.OrderHandlers.CreateOrder))
	r.Handle("GET", "/order/", http.HandlerFunc(h.OrderHandlers.OrderList))

	r.Handle("POST", "/courier/", http.HandlerFunc(h.OrderHandlers.CreateCourier))
	r.Handle("GET", "/courier/", http.HandlerFunc(h.OrderHandlers.CourierList))
	r.Handle("PUT", "/courier/{id}", http.HandlerFunc(h.OrderHandlers.UpdateCourier))
	r.Handle("DELETE", "/courier/{id}", http.HandlerFunc(h.OrderHandlers.DeleteCourier))
	r.Handle("GET", "/courier/{id}", http.HandlerFunc(h.OrderHandlers.GetCourier))

}
