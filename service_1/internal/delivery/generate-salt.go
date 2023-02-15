package delivery

import (
	"encoding/json"
	"net/http"
	"service_1/internal/models"
)

func (h *Handler) GenerateSalt(w http.ResponseWriter, r *http.Request) {

	resp := models.Response{
		Salt: h.service1.GenerateSalt(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
