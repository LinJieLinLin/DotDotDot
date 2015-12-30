package controllers

import (
	. "dot/models/common"
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type GetUser struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplNames = "index.html"
	return
}
func (c *GetUser) GetUser() {
	fmt.Println("hehe")
	re := Re{0, "123", "success"}
	c.Data["json"] = re
	c.ServeJson()
	return
}
