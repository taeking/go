package main

import (
	"fmt"
	"time"
)

func newtask() {
	for {
		fmt.Println("this is a new task")
		time.Sleep(time.Second)
	}

}

func main() {
	go newtask()
	for {
		fmt.Println("helloworld")
		time.Sleep(time.Second)
	}

}
