package transport

import (
	"encoding/json"
	"html"
	"net/http"
	"time"

	"log"
)

const (
	KeyContentType   = "Content-Type"
	KeyAuthorization = "Authorization"
	ContentTypeJSON  = "application/json; charset=utf-8"
	ContentTypeXML   = "application/xml; charset=utf-8"
	ContentTypeHTML  = "text/html; charset=utf-8"
)

// RESTFunc is a handler function that handles error and writes response in JSON
type RESTFunc func(*http.Request) (*Success, error)

// ServeHTTP implement http.Handler interface to write success or error response in JSON
func (h RESTFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get execution time
	start := time.Now()
	// Init http status
	var httpStatus int
	// Execute handler
	result, err := h(r)
	// If an error returned, return error
	if err != nil {
		httpStatus = sendErrorJSON(w, err)
	} else if result == nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		// if header exist, add header to response
		if result.Header != nil {
			for k, v := range result.Header {
				w.Header().Set(k, v)
			}
		}
		// send json success
		httpStatus = sendJSON(w, http.StatusOK, result)
	}
	// Log elapsed time
	log.Printf("HTTP Status: %d, Request: %s %s, Time elapsed: %s", httpStatus, r.Method,
		html.EscapeString(r.URL.Path), time.Since(start))
}

// sendJSON write response in JSON
func sendJSON(w http.ResponseWriter, httpStatus int, obj interface{}) int {
	// Add content type
	w.Header().Add(KeyContentType, ContentTypeJSON)
	// Write http status
	w.WriteHeader(httpStatus)
	// Send JSON response
	err := json.NewEncoder(w).Encode(obj)
	if err != nil {
		log.Println("Failed to write response", err)
	}
	// Return httpStatus
	return httpStatus
}

// sendErrorJSON write error response in JSON
func sendErrorJSON(w http.ResponseWriter, err error) int {
	// CastError error to Error
	apiError := CastError(err)
	// Show debug error response
	if apiError == ErrInternalServer {
		// Create response
		res := errorResponse{
			Code:    apiError.Code,
			Message: apiError.Message,
			Debug: &errorDebug{
				Message: err.Error(),
			},
		}
		return sendJSON(w, apiError.Status, &res)
	}
	// Send error json
	return sendJSON(w, apiError.Status, apiError)
}

// parseJSON parse json request body to o (target) and returns error
func ParseJSON(o interface{}, r *http.Request) error {
	d := json.NewDecoder(r.Body)
	if err := d.Decode(o); err != nil {
		return err
	}
	return nil
}
