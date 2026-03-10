package controllers

import (
	"bee-GO/models"
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
)

type BookController struct{
	web.Controller
}




func (c *BookController) GetAll() {
	c.Data["json"] = models.Books
	c.ServeJSON()
}


///GET BOOK BY ID
func (c  *BookController ) GetOne(){
	//now we have gotten the id from the url 
	id := c.Ctx.Input.Param(":id")
	//lets loop through the books one by one 
	for _, book := range models.Books {
		if book.ID == id {
			c.Data["json"] = book
			c.ServeJSON()
			return
		}
	}

	c.Data["json"] = "the book wasnt found"
	c.ServeJSON()
}


/////lets now create a book and add it to the slice  ,error handling included
//json


func (c *BookController) Create() {
	var book models.Book

	//converting the data into json ie parsing the request 
	err := json.Unmarshal(c.Ctx.Input.RequestBody,&book)
	if err != nil {
		///we send a status code of 400 if we hit this block or if there is an error 
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error:":"invalid json"}
		c.ServeJSON()
		return
	}
	if book.Author == "" || book.ID == ""  || book.Title == "" || book.Quantity == 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error:":"empty fields"}
		c.ServeJSON()
		return

	}
	models.Books = append(models.Books,book)
	c.Data["json"] = book
	c.ServeJSON()


}
















// func (c *BookController) Create() {
// 	var book models.Book

// 	// Unmarshal JSON from request body
// 	err := json.Unmarshal(c.Ctx.Input.RequestBody, &book)
// 	if err != nil {
// 		// Send 400 Bad Request with JSON error message
// 		c.Ctx.Output.SetStatus(400) // set HTTP status code
// 		c.Data["json"] = map[string]string{"error": "Invalid JSON"}
// 		c.ServeJSON()
// 		return
// 	}

// 	// Optional: basic validation
// 	if book.ID == "" || book.Title == "" || book.Author == "" {
// 		c.Ctx.Output.SetStatus(400)
// 		c.Data["json"] = map[string]string{"error": "Missing required fields"}
// 		c.ServeJSON()
// 		return
// 	}

// 	// Add book to the slice
// 	models.Books = append(models.Books, book)

// 	// Return the newly created book as JSON
// 	c.Data["json"] = book
// 	c.ServeJSON()
// }

