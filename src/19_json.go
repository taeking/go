package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string
	Subject []string
	Isok    bool
	Price   float64
}

func main() {
	it := IT{"santing", []string{"haha", "hello", "suzhou"}, true, 666.666}

	buf, err := json.Marshal(it)
	if err != nil {
		return
	}
	fmt.Println("buf=", string(buf))

}
