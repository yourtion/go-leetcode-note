package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	"leetcode-note/models"
)

type ExportController struct {
	beego.Controller
}

func (c *ExportController) Get() {
	// Session 验证
	user := c.GetSession(LoginKey)
	if !verifyUser(user) {
		c.Redirect("/", 307)
		return
	}

	c.TplName = "note/blog.md"

	s, e := getWeekRange()
	logs.Trace("GetWeek Range: %d -> %d", s, e)
	start, _ := c.GetInt("start", s)
	end, _ := c.GetInt("end", e)

	o := orm.NewOrm()
	qs := o.QueryTable("note").RelatedSel().OrderBy("Id")
	qs = qs.Filter("Day__gte", start)
	qs = qs.Filter("Day__lte", end)
	var notes []*models.Note
	if num, err := qs.All(&notes); err == nil {
		c.Data["count"] = num
		c.Data["data"] = &notes
		c.Data["start"] = start
		c.Data["end"] = end
	}
}
