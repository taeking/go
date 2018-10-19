package main

import (
	"fmt"

	"github.com/astaxie/goredis"
)

var client goredis.Client

func main() {
	client.Addr = "127.0.0.1:6379"
	err0 := client.Set("test_key", []byte("hello world"))
	if err0 != nil {
		fmt.Println("err0=", err0)
	}

	value, err := client.Get("test_key")
	if err != nil {
		fmt.Println("err=", err)
	} else {
		fmt.Println(string(value))
	}
}
