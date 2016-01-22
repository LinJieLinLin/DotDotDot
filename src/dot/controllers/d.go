package controllers

import (
	. "dot/models/common"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
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
	fmt.Println("SetName")
	re := Re{-1, "", ""}
	defer func() {
		c.Data["json"] = re
		c.ServeJson()
	}()
	u := Name{}
	valid := validation.Validation{}
	u.Name = c.GetString("name")
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
	_, err = o.Insert(&u)
	if nil != err {
		beego.Error("DB:", err)
		re.Msg = E1
		return
	}

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
	return
}
func (c *DCtrl) GetMenu() {
	fmt.Println("GetMenu")
	re := Re{-1, "", ""}

	defer func() {
		c.Data["json"] = re
		c.ServeJson()
	}()

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

	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("b.id",
		"n.name",
		"m.name as menu_name",
		"m.type",
		"m.price",
		"b.time",
	).
		From("tem_buy as b").
		InnerJoin("tem_menu as m").On("b.mid = m.id").
		InnerJoin("tem_name as n").On("b.uid = n.id").
		// Where("age > ?").
		OrderBy("id").Desc()
	// Limit(10).Offset(0)

	// 导出SQL语句
	sql := qb.String()

	qb, _ = orm.NewQueryBuilder("mysql")
	qb.Select(
		"name",
		"group_concat(menu_name separator ',') as menu_name ",
		"SUM(price) as price",
		"time",
	)
	qb.From(" (" + sql + ")" + " as list")
	qb.GroupBy("name")
	sql = qb.String()
	beego.Debug(sql)

	list := []*List{}
	// 执行SQL语句
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&list)

	buy := []*Buy{}
	num, err := o.QueryTable("tem_buy").All(&buy)
	if nil != err {
		beego.Error("DB:", err)
		re.Msg = E1
		return
	}
	if num != 0 {
		re.Code = 0
		re.Data = list
	}
	re.Msg = ""
	return
}
func (c *DCtrl) SetList() {
	fmt.Println("SetList")
	re := Re{-1, "", ""}
	defer func() {
		c.Data["json"] = re
		c.ServeJson()
	}()

	resData := Buy{}
	if err := json.Unmarshal([]byte(c.GetString("data")), &resData); err != nil {
		beego.Error(err)
		re.Msg = E0
		return
	}
	valid := validation.Validation{}
	b, err := valid.Valid(&resData)
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
	beego.Debug(resData)
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
	Id   int64  `form:"-"`
	Uid  int64  `form:"uid";valid:"Required"`
	Mid  int64  `form:"mid";valid:"Required"`
	Time string `form:"-"`
}
type Menu struct {
	Id    int64
	Name  string
	Type  int64
	Price float64
}
type Name struct {
	Id   int64
	Name string `valid:"Required";MaxSize(10)`
}
type List struct {
	Id       int64
	Name     string
	MenuName string
	Type     int64
	Price    float64
	Time     string
}

func init() {
	orm.RegisterModelWithPrefix("tem_", new(Buy))
	orm.RegisterModelWithPrefix("tem_", new(Menu))
	orm.RegisterModelWithPrefix("tem_", new(Name))
}
