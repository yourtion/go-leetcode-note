package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Layout = "layout.html"
	c.TplName = "index.tpl"

	c.Data["User"] = "Yourtion"
}
