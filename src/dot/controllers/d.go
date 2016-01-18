package controllers

import (
	. "dot/models/common"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
)

type DCtrl struct {
	beego.Controller
}

func (c *DCtrl) GetName() {
	fmt.Println("GetName")
	re := Re{-1, "", ""}

	defer func() {
		c.Data["json"] = re
		c.ServeJson()
	}()

	// u := Name{}
	// valid := validation.Validation{}
	// u.Account = c.GetString("account")
	// b, err := valid.Valid(&u)
	// if err != nil {
	//     beego.Error(err)
	//     re.Msg = E0
	//     return
	// }
	// if !b {
	//     for _, err := range valid.Errors {
	//         beego.Error(err.Key, err.Message)
	//         re.Msg = E0
	//         return
	//     }
	// }

	o := orm.NewOrm()
	name := []*Name{}
	num, err := o.QueryTable("tem_name").All(&name)
	if nil != err {
		beego.Error("DB:", err)
		re.Msg = E1
		return
	}
	if num != 0 {
		re.Code = 0
		re.Data = name
	}
	re.Msg = ""
	return
}

func (c *DCtrl) SetName() {
	fmt.Println("hehe")
	re := Re{0, "123", "success"}
	c.Data["json"] = re
	c.ServeJson()
	return
}
func (c *DCtrl) GetMenu() {
	fmt.Println("GetMenu")
	re := Re{-1, "", ""}

	defer func() {
		c.Data["json"] = re
		c.ServeJson()
	}()

	// u := Name{}
	// valid := validation.Validation{}
	// u.Account = c.GetString("account")
	// b, err := valid.Valid(&u)
	// if err != nil {
	//     beego.Error(err)
	//     re.Msg = E0
	//     return
	// }
	// if !b {
	//     for _, err := range valid.Errors {
	//         beego.Error(err.Key, err.Message)
	//         re.Msg = E0
	//         return
	//     }
	// }

	o := orm.NewOrm()
	menu := []*Menu{}
	num, err := o.QueryTable("tem_menu").All(&menu)
	if nil != err {
		beego.Error("DB:", err)
		re.Msg = E1
		return
	}
	if num != 0 {
		re.Code = 0
		re.Data = menu
	}
	re.Msg = ""
	return
}
func (c *DCtrl) SetMenu() {
	fmt.Println("hehe")
	re := Re{0, "123", "success"}
	c.Data["json"] = re
	c.ServeJson()
	return
}
func (c *DCtrl) GetList() {
	fmt.Println("GetList")
	re := Re{-1, "", ""}

	defer func() {
		c.Data["json"] = re
		c.ServeJson()
	}()

	// u := Name{}
	// valid := validation.Validation{}
	// u.Account = c.GetString("account")
	// b, err := valid.Valid(&u)
	// if err != nil {
	//     beego.Error(err)
	//     re.Msg = E0
	//     return
	// }
	// if !b {
	//     for _, err := range valid.Errors {
	//         beego.Error(err.Key, err.Message)
	//         re.Msg = E0
	//         return
	//     }
	// }

	o := orm.NewOrm()
	buy := []*Buy{}
	num, err := o.QueryTable("tem_buy").All(&buy)
	if nil != err {
		beego.Error("DB:", err)
		re.Msg = E1
		return
	}
	if num != 0 {
		re.Code = 0
		re.Data = buy
	}
	re.Msg = ""
	return
}

type Buy struct {
	Id   int64
	Uid  int64
	Mid  int64
	Time string
}
type Menu struct {
	Id    int64
	Name  string
	Type  int64
	Price float64
}
type Name struct {
	Id   int64
	Name string
}

func init() {
	orm.RegisterModelWithPrefix("tem_", new(Buy))
	orm.RegisterModelWithPrefix("tem_", new(Menu))
	orm.RegisterModelWithPrefix("tem_", new(Name))
}
