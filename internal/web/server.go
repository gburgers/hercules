// This is the package that sets up your web server, routes, and middleware.
package web

import (
	"log"
	"net/http"
	"your-project/internal/config"
)

func StartServer() {
	router := SetupRouter()
	log.Printf("Starting server on port %s", config.Cfg.Port)
	log.Fatal(http.ListenAndServe(":"+config.Cfg.Port, router))
}
