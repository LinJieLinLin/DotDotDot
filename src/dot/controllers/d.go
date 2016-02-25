package controllers

import (
	. "dot/models/common"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"time"
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
	menuList, msg := getMenu()
	if msg != "" {
		re.Msg = msg
		return
	}
	re.Code = 0
	re.Data = menuList
	re.Msg = ""
	return
}
func getMenu() ([]*Menu, string) {
	msg := ""
	o := orm.NewOrm()
	menuList := []*Menu{}
	num, err := o.QueryTable("tem_menu").All(&menuList)
	if nil != err {
		beego.Error("DB:", err)
		msg = E1
		return menuList, msg
	}
	beego.Info("menuList is", num)
	msg = ""
	return menuList, msg
}
func (c *DCtrl) SetMenu() {
	fmt.Println("SetMenu")
	re := Re{-1, "", ""}
	defer func() {
		c.Data["json"] = re
		c.ServeJson()
	}()
	menu := Menu{}
	valid := validation.Validation{}
	if err := json.Unmarshal([]byte(c.GetString("data")), &menu); err != nil {
		beego.Error(err)
		re.Msg = E0
		return
	}
	b, err := valid.Valid(&menu)
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
	if e := setMenuOne(menu); e != "" {
		re.Msg = e
		return
	}
	//菜单
	menuList, msg := getMenu()
	if msg != "" {
		re.Msg = msg
		return
	}
	re.Code = 0
	re.Data = menuList
	re.Msg = ""
	return
}
func setMenuOne(menu Menu) string {
	o := orm.NewOrm()
	_, err := o.Insert(&menu)
	if nil != err {
		beego.Error("DB:", err)
		return E1
	}
	return ""
}
func (c *DCtrl) GetList() {
	fmt.Println("GetList")
	re := Re{-1, "", ""}

	defer func() {
		c.Data["json"] = re
		c.ServeJson()
	}()

	sTime := c.GetString("time")
	eTime := sTime
	if sTime == "" {
		sTime = time.Now().Format("2006-01-02") + " 00:00:00"
		sTime = time.Now().Format("2006-01-02") + " 23:59:59"
	} else {
		sTime = sTime + " 00:00:00"
		eTime = eTime + " 23:59:59"
	}
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("b.id",
		"m.id as mid",
		"n.name",
		"m.name as menu_name",
		"m.type",
		"m.price",
		"b.time",
	).
		From("tem_buy as b").
		InnerJoin("tem_menu as m").On("b.mid = m.id").
		InnerJoin("tem_name as n").On("b.uid = n.id").
		Where("time between '" + sTime + "' and '" + eTime + "'").
		OrderBy("id").Desc()
	// Limit(10).Offset(0)

	// 导出SQL语句
	sql := ""
	detaultSql := qb.String()

	qb, _ = orm.NewQueryBuilder("mysql")
	qb.Select(
		"name",
		"group_concat(menu_name separator ',') as menu_name ",
		"SUM(price) as price",
		"time",
	)
	qb.From(" (" + detaultSql + ")" + " as list")
	qb.GroupBy("name")
	sql = qb.String()
	beego.Debug(sql, sTime)

	list := []*List{}
	// 读取下单列表执行SQL语句
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&list)
	if len(list) == 0 {
		re.Code = 0
		re.Msg = E2
		return
	}
	qb, _ = orm.NewQueryBuilder("mysql")
	qb.Select(
		"menu_name",
		"COUNT(id) as count",
		"SUM(price) as price",
	)
	qb.From(" (" + detaultSql + ")" + " as list")
	qb.GroupBy("mid")
	sql = qb.String()
	beego.Debug(sql, sTime)

	menuCount := []*MenuCount{}
	//读取菜单价格
	o.Raw(sql).QueryRows(&menuCount)
	reData := ReListData{}
	reData.List = list
	reData.Menu = menuCount
	re.Code = 0
	re.Data = reData
	re.Msg = ""
	return
}
func (c *DCtrl) SetListOne() {
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
	msg := setListOne(resData)
	if msg != "" {
		re.Msg = E1
		return
	}
	re.Code = 0
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

	resData := []*Buy{}
	beego.Info(c.GetString("data"))
	if err := json.Unmarshal([]byte(c.GetString("data")), &resData); err != nil {
		beego.Error(err)
		re.Msg = E0
		return
	}

	if len(resData) == 0 {
		beego.Error("传入数据为空")
		re.Msg = E0
		return
	}

	o := orm.NewOrm()
	nowDate := time.Now().Format("2006-01-02")
	delSql := "DELETE from tem_buy where uid = ? and time between ? and ? "
	res, err := o.Raw(delSql, resData[0].Uid, nowDate+" 00:00:00", nowDate+" 23:59:59").Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		beego.Info("删除数量: ", num)
	}
	if err != nil {
		beego.Error(err)
		re.Msg = E1
		return
	}
	for _, v := range resData {
		valid := validation.Validation{}
		beego.Info(v)
		b, err := valid.Valid(v)
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
	}
	for _, v := range resData {
		msg := setListOne(*v)
		if msg != "" {
			re.Msg = E1
			return
		}
	}
	re.Code = 0
	re.Msg = ""
	return
}
func setListOne(buy Buy) string {
	msg := ""
	o := orm.NewOrm()
	buy.Time = time.Now().Format("2006-01-02 15:04:05")
	_, err := o.Insert(&buy)
	if nil != err {
		beego.Error("DB:", err)
		msg = E1
	}
	return msg
}

type Buy struct {
	Id   int64  `form:"-"`
	Uid  int64  `valid:"Required";form:"uid"`
	Mid  int    `valid:"Min(1)";form:"mid"`
	Time string `form:"-"`
}
type Menu struct {
	Id    int64   `form:"-"`
	Name  string  `valid:"Required";form:"name"`
	Type  int64   `valid:"Required";form:"type"`
	Price float64 `valid:"Required";form:"price"`
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

type MenuCount struct {
	MenuName string
	Price    float64
	Count    int64
}

type ReListData struct {
	List []*List
	Menu []*MenuCount
}

func init() {
	orm.RegisterModelWithPrefix("tem_", new(Buy))
	orm.RegisterModelWithPrefix("tem_", new(Menu))
	orm.RegisterModelWithPrefix("tem_", new(Name))
}
