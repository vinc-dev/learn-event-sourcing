package transport

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// RegisterRoute is a function to register route.
type RegisterRoute func(router *Router)

// Router is the instance of our router.
type Router struct {
	*mux.Router
}

// HandleREST handles request and calls the corresponding handler
func (r *Router) HandleREST(path string, f RESTFunc) *mux.Route {
	return r.NewRoute().Path(path).Handler(f)
}

// NewRouter is used to init our router
func NewRouter(baseURL string, f RegisterRoute) *Router {
	r := Router{
		Router: mux.NewRouter(),
	}

	// Create route for base url
	a := Router{Router: r.PathPrefix(baseURL).Subrouter()}

	// Register routes
	f(&a)

	// Return routes
	return &r
}

// EnableCORS lets routes to support CORS
// TODO: Standardize what to be allowed
func EnableCORS(r *Router) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions})
	return handlers.CORS(originsOk, headersOk, methodsOk)(r)
}
