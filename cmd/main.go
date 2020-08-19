package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cast"
	"github.com/vinc-dev/learn-event-sourcing/config"
	transport "github.com/vinc-dev/learn-event-sourcing/transport/http/router"
)

func main() {
	// Init Routers
	port := ":" + cast.ToString(config.ConfServerPort)
	r := transport.NewRouter("", routes)
	// Init Handler
	var h http.Handler
	// Check CORS Config
	if cast.ToBool(config.ConfCORSEnabled) {
		log.Println("CORS Enabled")
		h = transport.EnableCORS(r)
	} else {
		log.Println("CORS Disabled")
		h = r
	}
	// Init boot time
	start := time.Now()
	// Print boot time elapsed
	log.Println(fmt.Sprintf("Boot Time: %s", time.Since(start)))
	// Start server and logs when router is error
	log.Println("error while listening to incoming request", http.ListenAndServe(port, h))
}
