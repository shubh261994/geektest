package routing

import (
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var router *chi.Mux
var onceRouter sync.Once
var staticPath string

func init() {
	router = chi.NewRouter()
	configureRouter()
}

func GetRouter() *chi.Mux {
	return router
}

func configureRouter() {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(DefaultLogger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Timeout(60 * time.Second))
}
