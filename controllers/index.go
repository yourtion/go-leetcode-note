package controllers

import (
	"github.com/astaxie/beego"
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
	if verifyUser(user) {
		c.SetSession(LoginKey, "yourtion")
	}
	if user == "logout" {
		c.DelSession("user")
	}

	login := c.GetSession("user")
	if verifyUser(login) {
		c.Data["login"] = true
	}

	// 列表
	o := orm.NewOrm()
	qs := o.QueryTable("note").RelatedSel().OrderBy("-Id")
	var notes []*models.Note
	if num, err := qs.All(&notes); err == nil {
		c.Data["count"] = num
		c.Data["data"] = &notes
	}
}
