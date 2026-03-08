package controllers

import (
	"bee-GO/models"

	"github.com/beego/beego/v2/server/web"
)

type BookController struct{
	web.Controller
}




func (c *BookController) GetAll() {
	c.Data["json"] = models.Books
	c.ServeJSON()
}



func (c  *BookController ) GetOne(){
	
	id :=  c.Ctx.Input.Param(":id")
	for _, book := range models.Books {
		if book.ID == id {
			c.Data["json"] = book
			c.ServeJSON()
			return
		}
		
	}

	c.Data["json"] = "book is not found"
	c.ServeJSON()
}


