package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

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
	// 列表
	o := orm.NewOrm()
	qs := o.QueryTable("note").RelatedSel().OrderBy("-Id").Limit(size, (page-1)*size)
	var notes []*models.Note
	if num, err := qs.All(&notes, "Id", "Problem", "Day", "Submissions", "Mark"); err == nil {
		c.Data["count"] = num
		c.Data["data"] = &notes
	}
}
