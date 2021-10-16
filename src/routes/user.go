package routes

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"go-users/src/users"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	id, _ := strconv.Atoi(getRequestParam("id", r).(string))
	response, err := users.GetById(id)
	if err != nil {
		log.Println(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrorResponse{
				Success: false,
				Message: "No user found for given ID",
			})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Success: false,
			Message: "Something went wrong",
		})
		return
	}

	json.NewEncoder(w).Encode(response)
}
