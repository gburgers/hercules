// Define your application routes here using http.ServeMux or a router like chi or gorilla/mux.
package web

import (
	"net/http"
	"your-project/internal/handlers"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", handlers.GetUserHandler)
	return mux
}
