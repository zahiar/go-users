package routes

import (
	"encoding/json"
	"net/http"
)

type Status struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func HandleStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	status := Status{
		Success: true,
		Message: "",
	}

	json.NewEncoder(w).Encode(status)
}
