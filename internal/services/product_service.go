package services

import (
	"context"
	"fmt"

	"github.com/gburgers/hercules/internal/database/repository"
	"github.com/gburgers/hercules/pkg/models"
)

type ProductService struct {
	repo *repository.ProductRepository
}

// NewProductService creates a new ProductService
func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// CreateProduct wraps repository logic and adds validation or other business logic
func (s *ProductService) CreateProduct(ctx context.Context, product *models.Product) error {
	// Add any business logic here, e.g., validation
	if product.Name == "" {
		return fmt.Errorf("product name cannot be empty")
	}
	if product.Price == 0 {
		return fmt.Errorf("price cannot be empty")
	}

	// Call the repository method to insert the product
	return s.repo.CreateProduct(ctx, product)
}

// GetProductByID handles business logic before fetching the product
func (s *ProductService) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid product ID")
	}

	// Call the repository to fetch the product
	product, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error fetching product: %w", err)
	}
	if product == nil {
		return nil, fmt.Errorf("product not found")
	}
	return product, nil
}

// GetProductByPrice
func (s *ProductService) GetProductByPrice(ctx context.Context, price float64) (*models.Product, error) {
	product, _ := s.repo.GetProductByPrice(ctx, price)
	return product, nil
}
