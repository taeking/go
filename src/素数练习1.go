package main

import "fmt"

func main() {
	var a, b int
	a = 1
A:
	for a < 100 {
		a++
		for b = 2; b < a; b++ {
			if a%b == 0 {
				goto A

			}
		}
		fmt.Println(a, "是素数")

	}

}
