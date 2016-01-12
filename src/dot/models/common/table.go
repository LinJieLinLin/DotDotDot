package common

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64
	Account  string
	Password string
	Name     string
	Tel      string
	Pic      string
	Email    string
	Role_id  int64
	Status   int64
	E_time   string
	C_time   string
	Last_ip  string
}
type UserR struct {
	Account  string
	Password string
	Name     string
	Tel      string
	Pic      string
	Email    string
	Role_id  int64
	Status   int64
	E_time   string
	C_time   string
	Last_ip  string
}

func init() {
	orm.RegisterModelWithPrefix("d_", new(User))
}
