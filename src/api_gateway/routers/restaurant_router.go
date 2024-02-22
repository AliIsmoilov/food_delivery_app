package routers

import (
	"monorepo/src/api_gateway/handlers"
	"monorepo/src/api_gateway/pkg/tracing"
	"net/http"
)

func RegisterRestaurantRoutes(r *tracing.TracedServeMux, h *handlers.Handlers) {
	r.Handle("POST", "/restaurant/", http.HandlerFunc(h.RestaurantHandlers.CreateRestaurant))
	r.Handle("GET", "/restaurant/", http.HandlerFunc(h.RestaurantHandlers.ListRestaurants))

	r.Handle("POST", "/category/", http.HandlerFunc(h.RestaurantHandlers.CreateCategory))
	r.Handle("GET", "/category/", http.HandlerFunc(h.RestaurantHandlers.ListCategories))

	r.Handle("POST", "/item/", http.HandlerFunc(h.RestaurantHandlers.CreateItem))
	r.Handle("GET", "/restaurant/items/", http.HandlerFunc(h.RestaurantHandlers.RestaurantItemsGroupedByCategory))
}
