package response

import (
	"encoding/json"
	"net/http"
)

const (
	DEFAULT_FORMAT_DEEP = 8
)

// basicResponse is a struct to contain default response message
type basicResponse struct {
	Message string `json:"message"`
}

// successResponse is a struct to contain default success response
type successResponse struct {
	Success bool `json:"success"`
}

// Json is a function to return json object with given data and statusCode
func Json(w http.ResponseWriter, data interface{}, statusCode int) {
	parseHeader(w, statusCode, "application/json")
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}

// Text is a function to return raw text and statusCode
func Text(w http.ResponseWriter, msg string, statusCode int) {
	Json(w, basicResponse{Message: msg}, statusCode)
}

// Success is a function to return success status and statusCode
func Success(w http.ResponseWriter, success bool, statusCode int) {
	Json(w, successResponse{Success: success}, statusCode)
}

// NotFound is a function to return 404 statusCode with empty body
func NotFound(w http.ResponseWriter) {
	Json(w, struct{}{}, http.StatusNotFound)
}

// parseHeader is a function to parse content data and add it to response header
func parseHeader(w http.ResponseWriter, statusCode int, contentType string) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
}
