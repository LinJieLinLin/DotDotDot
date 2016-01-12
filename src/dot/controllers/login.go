/*
登陆、注册、拿用户数据等与用户相关的操作
*/
package controllers

import (
	. "dot/models/common"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type LoginCtrl struct {
	beego.Controller
}
type UserData struct {
	Account string `valid:"Required;MaxSize(16)"`
	Name    string `valid:"MaxSize(16)"`
}

/*
传入参数：account="要查找的用户名"

检查用户是否存在,存在返回code:0，否则返回code:1
*/
func (c *LoginCtrl) CheckUser() {
	fmt.Println("CheckUser")
	re := Re{-1, "", ""}

	defer func() {
		c.Data["json"] = re
		c.ServeJson()
	}()

	u := UserData{}
	valid := validation.Validation{}
	u.Account = c.GetString("account")
	b, err := valid.Valid(&u)
	if err != nil {
		beego.Error(err)
		re.Msg = E0
		return
	}
	if !b {
		for _, err := range valid.Errors {
			beego.Error(err.Key, err.Message)
			re.Msg = E0
			return
		}
	}

	o := orm.NewOrm()
	count, err := o.QueryTable("d_user").Filter("Account", u.Account).Count()
	if nil != err {
		beego.Error("DB:", err)
		re.Msg = E1
		return
	}
	if count == 0 {
		re.Code = 0
	}
	re.Msg = ""
	return
}

/*
传入参数：account,password

登陆成功返回code:0,否则返回code:-1
*/
func (c *LoginCtrl) Login() {
	fmt.Println("Login")
	re := Re{-1, "", ""}

	defer func() {
		c.Data["json"] = re
		c.ServeJson()
	}()
	return
}

/*
传入参数：account,password

注册成功返回code:0,否则返回code:-1
*/
func (c *LoginCtrl) Register() {
	fmt.Println("Register")
	re := Re{-1, "", ""}

	defer func() {
		c.Data["json"] = re
		c.ServeJson()
	}()
	return
}
