package main

import (
	"fmt"
)

type People struct {
	id   int
	name string
	sex  byte //字符类型
}

func main() {
	var p People = People{1, "mike", 'm'}
	fmt.Print("people=", p)
	p1 := People{id: 1, name: "hello"}
	fmt.Print("p1=", p1)

}
