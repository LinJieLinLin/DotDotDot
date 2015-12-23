package routers

import (
	"dot/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.DirectoryIndex = true
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index.html", &controllers.MainController{})
	beego.Router("/getUser", &controllers.GetUser{},"*:GetUser")
	beego.SetStaticPath("/dot", "www")
	beego.ViewsPath = "view"
}
