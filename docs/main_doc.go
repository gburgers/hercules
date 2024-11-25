package _

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	Name      string
	Price     float64
	Available bool
}

func main() {
	connStr := "postgres://postgres:secret@localhost:5432/gopgtest?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// create a table
	createProductTable(db)

	// create a product
	product := Product{"Book", 15.55, true}
	pk := insertProduct(db, product)

	// query a product from product table
	var name string
	var available bool
	var price float64

	query := "SELECT name, available, price FROM product WHERE id = $1"
	err = db.QueryRow(query, pk).Scan(&name, &available, &price)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No rows found with ID %d", pk)
		}
		log.Fatal(err)
	}
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Available: %t\n", available) // %t - prints true or false
	fmt.Printf("Price: %f\n", price)
}

func CreateProductTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price NUMERIC(6,2) NOT NULL,
    available BOOLEAN,
    created timestamp DEFAULT NOW()
  )`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func insertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO product (name, price, available)
    VALUES ($1, $2, $3) RETURNING id`

	var pk int
	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}
