package main

import (
	"fmt"
	"os"
)

func WriteFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer file.Close()
	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i=%d\n", i)
		file.WriteString(buf)
	}

}

func ReadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer file.Close()
	buf := make([]byte, 1024*2)
	n, err := file.Read(buf)
	fmt.Println("n=", string(buf[:n]))

}

func main() {

	path := "./demo.txt"
	//WriteFile(path)
	ReadFile(path)

}
