package main

import (
	"fmt"

	"github.com/astaxie/beego/httplib"
)

func main() {
	sUrl := "https://movie.douban.com/subject/25827935/"
	rep := httplib.Get(sUrl)
	url, err := rep.String()
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("url=", url)
}
