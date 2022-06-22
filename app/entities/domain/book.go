package domain

// Books is a slice of Book
type Books []Book

// book represents a book in the system.
type Book struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
