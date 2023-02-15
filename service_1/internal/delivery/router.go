package delivery

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(Mux *chi.Mux, Handler *Handler) {

	// Options
	Mux.Use(middleware.Logger)
	Mux.Use(middleware.Timeout(60 * time.Second))

	// healthcheck
	Mux.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("health"))
	})

	Mux.Post("/generate-salt", Handler.GenerateSalt)
}
