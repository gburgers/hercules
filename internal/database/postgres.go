// This package is responsible for setting up and managing the
// connection to the PostgreSQL database using pgxpool or gorm.
package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

func Connect() error {
	connStr := "postgres://postgres:secret@localhost:5432/gopgtest"

	// Use pgx as the sql driver
	var err error
	db, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
		return err
	}
	log.Println("Successfully connected to the PostgreSQL database using pgx!")
	return nil
}

func Close() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing the database connection: %v", err)
		} else {
			log.Println("Database connection closed")
		}
	}
}

// Function for returning *sql.DB for checking DB connection in other packages
func DB() *sql.DB {
	if db == nil {
		log.Println("DB() returning nil")
	}

	return db
}

// Function to check if the database is open
func CheckDatabaseConnection() {
	// db := database.DB()
	if db != nil {
		err := db.Ping()
		if err != nil {
			fmt.Println("Database connection is NOT open:", err)
		} else {
			fmt.Println("Database connection is open and healthy.")
		}
	} else {
		fmt.Println("Database connection is not established.")
	}
}
