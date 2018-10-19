package main

import "fmt"

func main() {
	var a int = 10
	var p *int
	p = &a

	fmt.Println("a=", a)
	fmt.Println("p=", p)
	fmt.Printf("a=%d\n", a)
	fmt.Printf("p=%T\n", p)
	fmt.Printf("*p=%v\n", *p)
	fmt.Print("---------------------\n")

	*p = 1000000

	fmt.Printf("*p=%v\n", *p)

}
