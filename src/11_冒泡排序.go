package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var a [10]int

	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(100)
		fmt.Printf("%d,", a[i])
	}

	for j := 0; j < len(a)-1; j++ {
		for k := 0; k < len(a)-1-j; k++ {
			if a[k] > a[k+1] {
				a[k], a[k+1] = a[k+1], a[k]
			}

		}

	}
	fmt.Println("排序后")
	for i := 0; i < len(a); i++ {

		fmt.Printf("%d,", a[i])
	}

}
