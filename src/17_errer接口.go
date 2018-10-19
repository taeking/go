package main

import "fmt"
import "errors"

func main() {

	var err1 error = fmt.Errorf("%s", "this is normal error")
	fmt.Println("err1=", err1)

	err2 := errors.New("this is normal")
	fmt.Println("err2=", err2)

}
