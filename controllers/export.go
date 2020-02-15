package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

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

	err, notes, num := models.GetNotesBetweenStartAndEnd(start, end)
	if err != nil {
		return
	}
	c.Data["count"] = num
	c.Data["data"] = &notes
	c.Data["start"] = start
	c.Data["end"] = end
}
