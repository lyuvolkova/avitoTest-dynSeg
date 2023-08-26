package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func sendResponse(response interface{}, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	statusCode := http.StatusOK

	if err != nil {
		statusCode = http.StatusInternalServerError

		switch {
		case errors.Is(err, ErrBadRequest):
			statusCode = http.StatusBadRequest
		case errors.Is(err, ErrNotImplemented):
			statusCode = http.StatusNotImplemented
		}

		response = ErrorResponse{
			Error: err.Error(),
		}
	}

	w.WriteHeader(statusCode)

	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
	}
}

func PathVarInt64(r *http.Request, name string) (int64, error) {
	value := mux.Vars(r)[name]

	return strconv.ParseInt(value, 10, 64)
}
