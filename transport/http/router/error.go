package transport

import (
	"log"
	"net/http"
)

// Standard Errors
var ErrBadRequest = Error{
	Status:  http.StatusBadRequest,
	Code:    "400",
	Message: "Bad Request",
}
var ErrUnauthorized = Error{
	Status:  http.StatusUnauthorized,
	Code:    "401",
	Message: "Unauthorized",
}
var ErrForbidden = Error{
	Status:  http.StatusForbidden,
	Code:    "403",
	Message: "Forbidden",
}
var ErrNotFound = Error{
	Status:  http.StatusNotFound,
	Code:    "404",
	Message: "Not Found",
}
var ErrMethodNotAllowed = Error{
	Status:  http.StatusMethodNotAllowed,
	Code:    "405",
	Message: "Method Not Allowed",
}
var ErrInternalServer = Error{
	Status:  http.StatusInternalServerError,
	Code:    "500",
	Message: "Internal Error",
}

// Error represent standard API error that contains HTTP Status (Status) and API-scoped Error Code (Code).
type Error struct {
	Status  int    `yaml:"status" json:"-"`
	Code    string `yaml:"code" json:"code"`
	Message string `yaml:"message" json:"message"`
}

// Error is an implementation of built-in error type interface
func (e Error) Error() string {
	return e.Message
}

// CastError cast error interface as an Error
func CastError(err error) Error {
	apiErr, ok := err.(Error)
	if !ok {
		log.Println(err)
		// If assert type fail, create new internal error
		apiErr = ErrInternalServer
	}
	return apiErr
}
