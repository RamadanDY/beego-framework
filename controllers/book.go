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


