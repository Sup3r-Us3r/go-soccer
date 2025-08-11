package apperr

import (
	"encoding/json"
	"net/http"
)

type AppErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (ae *AppErr) Error() string {
	return ae.Message
}

func NewAppErr(message string, err string, code int) *AppErr {
	return &AppErr{
		Code:    code,
		Message: message,
	}
}

func NewHttpError(w http.ResponseWriter, errorData *AppErr) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorData.Code)
	json.NewEncoder(w).Encode(errorData)
}

func NewBadRequestError(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewUnauthorizedError(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func NewForbiddenError(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusForbidden,
		Message: message,
	}
}

func NewNotFoundError(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewInternalServerError(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
