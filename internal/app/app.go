package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gburgers/hercules/internal/database"
	"github.com/gburgers/hercules/internal/database/repository"
	"github.com/gburgers/hercules/internal/services"
	"github.com/gburgers/hercules/pkg/models"
)

// App contains all dependencies for the application.
type App struct {
	ProductService *services.ProductService
}

// InitializeApp sets up the application (database, services, etc.)
func InitializeApp() (*App, error) {
	// Connect to the database
	err := database.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	fmt.Println("Application running...")
	database.CheckDatabaseConnection()

	// Create a new UserRepository instance
	productRepo := repository.NewProductRepository(database.DB())

	// Example: Creating a new user
	ctx := context.Background()
	newProduct := &models.Product{Name: "Vitamin_A", Price: 12.95}
	err = productRepo.CreateProduct(ctx, newProduct)
	if err != nil {
		log.Fatalf("Error creating product: %v", err)
	}
	log.Println("Product created successfully")

	// Create service instances
	productService := services.NewProductService(productRepo)

	return &App{
		ProductService: productService,
	}, nil
}

// StartServer (optional) starts the HTTP server
func (app *App) StartServer() {
	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		// Example endpoint to fetch product
		// Delegate to ProductService
		product, err := app.ProductService.GetProductByPrice(r.Context(), 12.95)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Product: %v\n", product)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// CloseApp cleans up resources before the app shuts down
func (app *App) CloseApp() {
	log.Println("Shutting down the application and closing the database.")
	database.Close() // Clean up the database connection
}
