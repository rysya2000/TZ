package app

import (
	"fmt"
	"log"
	"net/http"
	"service_1/config"
	"service_1/internal/delivery"
	"service_1/internal/usecase"

	"github.com/go-chi/chi"
)

// Run ...
func Run(cfg *config.Config) {
	fmt.Println(cfg)

	usecase := usecase.New()
	handler := delivery.New(usecase)

	//  HTTP Server
	Mux := chi.NewRouter()
	delivery.NewRouter(Mux, handler)

	log.Printf("Starting up on http://localhost:%s", cfg.HTTP.Port)

	log.Fatal(http.ListenAndServe(":"+cfg.HTTP.Port, Mux))
}
