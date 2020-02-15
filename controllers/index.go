package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"leetcode-note/models"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Layout = "layout.html"
	c.TplName = "index.tpl"

	// Session
	user := c.Ctx.Input.Query(LoginKey)
	logs.Trace("user: %v", user)
	if verifyUser(user) {
		c.SetSession(LoginKey, user)
	}
	if user == "logout" {
		c.DelSession(LoginKey)
	}

	login := c.GetSession(LoginKey)
	if verifyUser(login) {
		c.Data["login"] = true
	}

	page, _ := c.GetInt("page", 1)
	size, _ := c.GetInt("size", 20)
	logs.Trace("Page Info: %d , %d", page, size)
	err, notes, num := models.PagesNotes(page, size)
	if err != nil {
		return
	}
	c.Data["count"] = num
	c.Data["data"] = &notes
}
