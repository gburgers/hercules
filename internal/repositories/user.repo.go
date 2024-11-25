// Repositories handle all the database access logic (CRUD operations).
// These abstract away raw SQL queries or ORM operations, keeping your services and handlers clean.
package repositories

import (
	"context"
	"your-project/internal/database"
	"your-project/internal/models"
)

func GetUserByID(ctx context.Context, id int) (*models.User, error) {
	query := `SELECT id, name, email, password FROM users WHERE id = $1`
	row := database.DbPool.QueryRow(ctx, query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
