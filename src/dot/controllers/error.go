package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

//404页面
func (c *ErrorController) Error404() {
	c.Data["content"] = "page not found"
	c.TplNames = "404.html"
}

//501页面
func (c *ErrorController) Error501() {
	c.Data["content"] = "server error"
	c.TplNames = "501.tpl"
}

//数据库连接失败的页面
func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplNames = "dberror.tpl"
}
