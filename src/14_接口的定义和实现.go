package main

import (
	"fmt"
)

type Humaner interface {
	sayhi()
}
type Student struct {
	name string
	id   int
}

func (tmp *Student) sayhi() {
	fmt.Printf("Student[%s,%d]\n", tmp.name, tmp.id)

}

type Teacher struct {
	addr   string
	grouop string
}

func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher[%s,%s]\n", tmp.addr, tmp.grouop)

}

type Mystr string

func (tmp *Mystr) sayhi() {
	fmt.Printf("Mystr[%s]\n", *tmp)

}
func main() {
	var i Humaner
	j := &Student{"mike", 666}
	i = j
	i.sayhi()

	k := &Teacher{"hehe", "suzhou"}
	i = k
	i.sayhi()

	var m Mystr = "helloword"
	i = &m
	i.sayhi()

}
