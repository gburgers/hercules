// This package contains business logic. It interacts with repositories and handles
// more complex operations like validation, transactions, or combining multiple data sources.
package services

import (
	"context"
	"your-project/internal/models"
	"your-project/internal/repositories"
)

func GetUserProfile(ctx context.Context, userID int) (*models.User, error) {
	return repositories.GetUserByID(ctx, userID)
}
