package delivery

import (
	"service_1/internal/usecase"
)

type Handler struct {
	service1 usecase.Service_1
}

func New(service1 usecase.Service_1) *Handler {
	return &Handler{
		service1: service1,
	}
}
