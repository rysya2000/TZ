package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"service_2/internal/models"
	"service_2/internal/usecase/utils"
)

var postURL = "http://172.1.23.12:8080/generate-salt"

type User struct {
	repo CreateAndRead
}

func NewUser(repo CreateAndRead) *User {
	return &User{
		repo: repo,
	}
}

func (r *User) CreateUser(u models.User) error {
	if u.ValidateEmail() == false {
		return utils.ErrInvalidEmail
	}

	_, err := r.repo.ReadUser(u)
	if err == nil {
		return utils.ErrAlreadyRegistered
	}

	req, err := http.NewRequest(http.MethodPost, postURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return err
	}

	us := models.User{}

	err = json.NewDecoder(res.Body).Decode(&us)
	if err != nil {
		return err
	}

	u.Salt = us.Salt
	u.Password = Hash(us.Salt, u.Password)

	return r.repo.CreateUser(u)
}

func (r *User) ReadUser(u models.User) (models.User, error) {
	// fmt.Printf("initial user: %v\n", u)
	user, err := r.repo.ReadUser(u)
	if err != nil {
		return models.User{}, utils.ErrUnathorized
	}

	// fmt.Printf("Read USER - user: %v\n\n", user)
	return user, nil
}
