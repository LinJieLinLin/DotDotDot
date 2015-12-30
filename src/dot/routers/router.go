package routers

import (
	"dot/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.DirectoryIndex = true
	beego.TemplateLeft = "<<<"
	beego.TemplateRight = ">>>"
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index.html", &controllers.MainController{})
	beego.Router("/getUser", &controllers.GetUser{}, "*:GetUser")
	beego.SetStaticPath("/dot", "www")
	beego.Router("/*", &controllers.ErrorController{}, "*:Error404")
	beego.ViewsPath = "view"
}
