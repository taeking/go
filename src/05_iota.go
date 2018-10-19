package main

import "fmt"

func main() {

	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
	const d = iota
	fmt.Printf("d=%d", d)

}
