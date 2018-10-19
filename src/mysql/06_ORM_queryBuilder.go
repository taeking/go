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
	orm.Debug = true
	var users []User

	orm.RegisterDataBase("default", "mysql", "root:123456@/test?charset=utf8")
	orm.RegisterModel(new(User))
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("name").From("user").Where("id=?").Limit(1)
	sql := qb.String()
	o.Raw(sql, "5").QueryRows(&users)
	fmt.Printf("结果是：%v\n", users)

}
