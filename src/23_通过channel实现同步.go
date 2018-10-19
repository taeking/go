package main

import (
	"fmt"
	"time"
)

var cha = make(chan int)

func printer(str string) {
	for _, data := range str {
		fmt.Print(string(data))
		time.Sleep(time.Second)
	}
}

func person1() {
	printer("helloworld")
	cha <- 666

}
func person2() {
	<-cha
	printer("helloworld")

}

func main() {

	go person1()
	go person2()
	for {

	}

}
