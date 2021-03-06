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
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	user := new(User)
	user.Id = 1
	o.Read(user)

	fmt.Println(user)

}
