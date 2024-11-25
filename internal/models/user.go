// This contains Go structs that represent the data in your PostgreSQL tables.
// These structs map to the columns in your database.
package models

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}
