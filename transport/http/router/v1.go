package transport

func V1(r *Router) {
	//eventService := eventsource.NewService()

	r.HandleREST("v1/book", nil).Methods("POST")
}
