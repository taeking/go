package main

import (
	"fmt"
)

type Student struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main() {
	var s1 Student = Student{1, "mike", 'm', 18, "beijing"}
	fmt.Println("s1=", s1)

	s2 := Student{name: "mike", age: 18}
	fmt.Println("s2=", s2)

}
