package usecase

import (
	"service_2/internal/models"
)

type (
	CreateAndRead interface {
		CreateUser(models.User) error
		ReadUser(models.User) (models.User, error)
	}
)
