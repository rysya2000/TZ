package delivery

import (
	"service_2/internal/models"
	"service_2/internal/usecase/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = h.usecase.User.CreateUser(req)
	if err != nil {
		if err == utils.ErrAlreadyRegistered {
			http.Error(w, err.Error(), 400)
		} else if err == utils.ErrInvalidEmail {
			http.Error(w, err.Error(), 400)
		} else {
			http.Error(w, err.Error(), 500)
		}
		return
	}

	w.WriteHeader(200)

}

func (h *Handler) ReadUser(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	u := models.User{
		Email: email,
	}

	user, err := h.usecase.User.ReadUser(u)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	fmt.Println(user)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
