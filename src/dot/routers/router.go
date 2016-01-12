package routers

import (
	. "dot/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//db
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@/dot?charset=utf8")
	//set
	beego.DirectoryIndex = true
	beego.TemplateLeft = "<<<"
	beego.TemplateRight = ">>>"
	beego.ErrorController(&ErrorController{})
	beego.Router("/", &MainController{})

	//login
	beego.Router("/n/checkUser", &LoginCtrl{}, "get:CheckUser")
	beego.Router("/index.html", &MainController{})
	beego.Router("/getUser", &GetUser{}, "*:GetUser")
	beego.SetStaticPath("/dot", "www")
	beego.Router("/*", &ErrorController{}, "*:Error404")
	beego.ViewsPath = "view"
}
