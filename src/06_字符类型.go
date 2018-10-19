package main

import "fmt"

func main() {
	var ch byte
	ch = 97
	//%c以字符方式打印，%d以整型方式打印
	fmt.Printf("%c\n", ch)
	fmt.Printf("%d\n", ch)
	var aa byte
	aa = 'a'
	var bb string
	bb = "a"
	fmt.Println(aa)
	fmt.Println(bb)

}
