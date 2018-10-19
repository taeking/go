package main

import (
	"database/sql"
	"fmt"

	//"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("err=", err)
	}
	stmt, err1 := db.Prepare("INSERT user SET name=?,sex=?")
	if err1 != nil {
		fmt.Println("err1=", err1)
	}
	_, err2 := stmt.Exec("zhangsan", 1)
	if err2 != nil {
		fmt.Println("err2=", err2)
	}

}
