package main

import "fmt"

func main() {
	var id [50]int
	for i := 0; i < len(id); i++ {
		id[i] = i + 1
		fmt.Printf("id[%d]=%d\n", i, id[i])

	}
	for i, data := range id {
		fmt.Printf("id[%d]=%d\n", i, data)

	}

}
