package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	coon, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("coon err", err)
		return
	}
	defer coon.Close()
	_, err1 := coon.Do("SET", "aaa", 100)
	if err1 != nil {
		fmt.Println("SET err1", err1)
		return
	}
	r, err2 := redis.Int(coon.Do("GET", "aaa"))
	if err2 != nil {
		fmt.Println("SET err2", err2)
		return
	}

	fmt.Print(r)

}
