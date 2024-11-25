package main

import (
	"fmt"
	"log"

	"github.com/gburgers/hercules/internal/database"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// connect to the database
	err := database.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	defer database.Close()

	fmt.Println("Application running...")
	database.CheckDatabaseConnection()
}
