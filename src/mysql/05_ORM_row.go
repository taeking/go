package main

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

func main() {

	orm.RegisterDataBase("default", "mysql", "root:123456@/test?charset=utf8")
	orm.RegisterModel(new(User))
	o := orm.NewOrm()
	o.Using("default")
	//var maps []orm.Params
	//o.Raw("select * from user").Values(&maps)

	var users []User
	o.Raw("select * from user").QueryRows(&users)

	fmt.Printf("结果是：%v\n", users)
	//for _, v := range maps {

	//	fmt.Printf("结果是：%v\n", v)
	//}
}
