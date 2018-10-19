package main

import "fmt"

func testa() {
	fmt.Println("aaaaa")
}
func testb() {

	panic("this is a  panic")
	//fmt.Println("bbbbb")
}
func testc() {
	fmt.Println("ccccc")
}

func main() {
	testa()
	testb()
	testc()

}
