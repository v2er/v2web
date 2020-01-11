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

func (c *IndexController) Topic() {
	id, _ := c.GetInt64(":id")
	page, _ := c.GetInt("p")

	data, err := c.client.Content(id, page)
	c.OnError(err)
	c.Data["data"] = data
	c.Tpl("topic.html", data.Topic.Title)
}
