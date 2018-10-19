package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func CheckError(err error) {
	fmt.Println("err=", err)
}

func main() {
	db, err1 := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
	CheckError(err1)
	stmt, err2 := db.Prepare("UPDATE user SET name=? where id=?")
	CheckError(err2)
	stmt.Exec("lisi", 3)

}
