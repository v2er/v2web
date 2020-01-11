package main

import (
	"regexp"
	_ "v2web/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.AddFuncMap("id", id)

	beego.Run()
}

func id(link string) string {
	res := regexp.MustCompile(`\/t\/(\d+)`).FindStringSubmatch(link)
	if len(res) > 1 {
		return res[1]
	}
	return ""
}
