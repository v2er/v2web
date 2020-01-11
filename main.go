package main

import (
	"html/template"
	"regexp"
	_ "v2web/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.AddFuncMap("id", id)
	beego.AddFuncMap("html", html)
	beego.AddFuncMap("page", page)

	beego.Run()
}

func id(link string) string {
	res := regexp.MustCompile(`\/t\/(\d+)`).FindStringSubmatch(link)
	if len(res) > 1 {
		return res[1]
	}
	return ""
}

func html(s string) template.HTML {
	return template.HTML(s)
}

func page(n int) (arr []int) {
	for i := 0; i < n; i++ {
		arr = append(arr, i+1)
	}
	return
}
