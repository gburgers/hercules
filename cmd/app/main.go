package main

import (
	"log"

	"github.com/gburgers/hercules/internal/app"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// Start the application
	application, err := app.InitializeApp()
	if err != nil {
		log.Fatalf("Error initializing the app: %v", err)
	}

	// Defer the closing of the app
	defer application.CloseApp()

	// Start the web server (if applicable) or other services
	application.StartServer() // connect to the database
}
