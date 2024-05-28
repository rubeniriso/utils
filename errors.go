package utils

import (
	"net/http"
)

type ApiError struct {
	Error string `json:"error"`
}

func permissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
}
func badRequest(w http.ResponseWriter) {
	WriteJSON(w, http.StatusForbidden, ApiError{Error: "bad request"})
}
