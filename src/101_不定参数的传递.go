package main

import "fmt"

func main() {
	testa(1, 2, 3, 4)
	cc := testc()
	fmt.Println(cc)

}
func testa(args ...int) {
	testb(args...)

}
func testc() (result int) {
	result = 666
	return

}
func testb(args ...int) {
	for _, data := range args {
		fmt.Println("data=", data)
	}

}
