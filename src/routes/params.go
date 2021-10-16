package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func getRequestParam(name string, r *http.Request) interface{} {
	params := mux.Vars(r)
	return params[name]
}
