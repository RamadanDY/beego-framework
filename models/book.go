package models


type Book struct {
	ID string `json:"id"`
	Title string	`json:"title"`
	Author string	`json:"author"`
	Quantity int	`json:"quantity"`
}

//lets create a slice for the list of books based on the struct above

var Books = []Book {
	{ID: "1", Title: "The Hobbit", Author: "Tolkien", Quantity: 5},
	{ID: "2", Title: "Harry Potter", Author: "Rowling", Quantity: 7},
}