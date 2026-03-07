package routers

import (
	"bee-GO/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
    web.Router("/", &controllers.MainController{})
	web.Router("/books", &controllers.BookController{}, "get:GetAll")

}
