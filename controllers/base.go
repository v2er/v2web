package controllers

import (
	"github.com/astaxie/beego"
	"github.com/v2er/v2api"
)

type BaseController struct {
	beego.Controller
	client *v2api.Client
}

func (c *BaseController) Prepare() {
	c.client = v2api.NewClient()
}

func (c *BaseController) Tpl(name, title string) {
	c.Layout = "layout.html"
	c.TplName = name
	c.Data["title"] = title
}

func (c *BaseController) OnError(err error) {
	if err != nil {
		c.Ctx.WriteString(err.Error())
		c.StopRun()
	}
}
