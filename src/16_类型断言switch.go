package main

import (
	"fmt"
)

type Student struct {
	name string
	id   int
}

func main() {

	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hell go"
	i[2] = Student{"milk", 666}
	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d]类型为int,内容为%d\n", index, value)
		case string:
			fmt.Printf("x[%d]类型为string,内容为%s\n", index, value)
		case Student:
			fmt.Printf("x[%d]类型为Student,内容为name=%s,id=%d\n", index, value.name, value.id)

		}

	}

}
