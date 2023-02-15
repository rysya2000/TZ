package utils

import "errors"

var (
	ErrAlreadyRegistered = errors.New("Already registered")
	ErrInvalidEmail      = errors.New("Invalid email address")
	ErrUnathorized       = errors.New("Unathorized")
)
