package main

import (
	"fmt"
)

func main() {
	var cha = make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			cha <- i

		}
		close(cha)
	}()

	for {
		if value, ok := <-cha; ok == true {
			fmt.Println("value=", value)
		}
	}
}
