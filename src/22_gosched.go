package main

import (
	"fmt"
	"runtime"
)

func newtask() {
	for i := 1; i < 5; i++ {

		fmt.Println("this is a new task")

	}

}

func main() {
	go newtask()
	runtime.Gosched()

	for i := 1; i < 2; i++ {

		fmt.Println("helloworld")

	}

}
