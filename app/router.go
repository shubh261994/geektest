package app

import (
	"github.com/go-chi/chi"

	"geektest/app/common/middleware"
	"geektest/app/orders"
)

func SetRoutes(router *chi.Mux) {
	router.Use(middleware.OptionsMiddleware)
	router.Get("/orders", orders.GetOrders)
}
