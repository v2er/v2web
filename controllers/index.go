package controllers

type IndexController struct {
	BaseController
}

func (c *IndexController) Home() {
	topics, err := c.client.Latest()
	c.OnError(err)
	c.Data["topics"] = topics
	c.Tpl("home.html", "V2EX")
}
