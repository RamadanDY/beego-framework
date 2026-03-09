package models


///this  is to enable  json and form tags to work 
type Book struct {
	ID       string `json:"id" form:"id"`
	Title    string `json:"title" form:"title"`
	Author   string `json:"author" form:"author"`
	Quantity int    `json:"quantity" form:"quantity"`
}


//but this is just for json tags 

// type Book struct {
// 	ID string `json:"id"`
// 	Title string	`json:"title"`
// 	Author string	`json:"author"`
// 	Quantity int	`json:"quantity"`
// }

//lets create a slice for the list of books based on the struct above

var Books = []Book {
	{ID: "1", Title: "The Hobbit", Author: "Tolkien", Quantity: 5},
	{ID: "2", Title: "Harry Potter", Author: "Rowling", Quantity: 7},
}