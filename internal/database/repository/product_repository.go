package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gburgers/hercules/pkg/models"
)

type ProductRepository struct {
	db *sql.DB
}

// NewProductRepository creates a new ProductRepository
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// CreateProduct inserts a new product into the database
func (r *ProductRepository) CreateProduct(ctx context.Context, product *models.Product) error {
	query := `INSERT INTO product (name, price) VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, product.Name, product.Price)
	if err != nil {
		return fmt.Errorf("could not insert product: %w", err)
	}
	return nil
}

// GetProductByID retrieves a product by their ID
func (r *ProductRepository) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
	query := `SELECT id, name, price product WHERE id = $1`
	product := &models.Product{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&product.ID, &product.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("could not fetch user: %w", err)
	}
	return product, nil
}

// UpdateProduct updates a Product's information
func (r ProductRepository) UpdateProduct(ctx context.Context, product *models.Product) error {
	query := `UPDATE product SET name=$1, price=$2 WHERE id=$3`
	_, err := r.db.ExecContext(ctx, query, product.Name, product.Price)
	if err != nil {
		return fmt.Errorf("could not update product: %w", err)
	}
	return nil
}

// DeleteProduct deletes a product by ID
func (r *ProductRepository) DeleteProduct(ctx context.Context, id int) error {
	query := `DELETE FROM product WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("could not delete product: %w", err)
	}
	return nil
}

// GetProductByPrice retrieves a product by their price
func (r *ProductRepository) GetProductByPrice(ctx context.Context, price float64) (*models.Product, error) {
	query := `SELECT id, name, price product WHERE price = $2`
	product := &models.Product{}
	err := r.db.QueryRowContext(ctx, query, price).Scan(&product.Price, &product.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("could not fetch user: %w", err)
	}
	return product, nil
}
