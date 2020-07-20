package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
	Sex  int
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:199411@tcp(192.168.31.205:3306)/test?charset=utf8")
	orm.RegisterModel(new(User))
}

func PrintUserByOrm() {
	o := orm.NewOrm()
	user := User{Id: 3}
	o.Read(&user)
	fmt.Println(user)
}
