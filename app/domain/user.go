package domain

// Users is a slice of User
type Users []User

// User represents a user in the system.
type User struct {
	ID    int64
	Name  string
	Email string
}
