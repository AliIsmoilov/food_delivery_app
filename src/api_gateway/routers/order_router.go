package routers

import (
	"monorepo/src/api_gateway/handlers"
	"monorepo/src/api_gateway/pkg/tracing"
	"net/http"
)

func RegisterOrderRoutes(r *tracing.TracedServeMux, h *handlers.Handlers) {
	r.Handle("POST", "/order/", http.HandlerFunc(h.OrderHandlers.CreateOrder))
	r.Handle("GET", "/order/", http.HandlerFunc(h.OrderHandlers.OrderList))
}
