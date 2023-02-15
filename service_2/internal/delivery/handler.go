package delivery

import (
	"service_2/internal/usecase"
)

type Handler struct {
	usecase usecase.UseCase
}

func New(uc usecase.UseCase) *Handler {
	return &Handler{
		usecase: uc,
	}
}
