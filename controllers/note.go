package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	"leetcode-note/models"
)

type NoteController struct {
	beego.Controller
}

func (c *NoteController) Get() {
	// Session 验证
	user := c.GetSession(LoginKey)
	if !verifyUser(user) {
		c.Redirect("/", 307)
		return
	}

	sid := c.Ctx.Input.Param(":id")
	if id, err := strconv.Atoi(sid); err == nil || id > 1 {
		if err, note := models.GetNoteById(id); err == nil {
			c.Data["note"] = &note
		}
	}
	c.Layout = "layout.html"
	c.TplName = "note/form.tpl"
}

func (c *NoteController) Post() {
	// Session 验证
	user := c.GetSession(LoginKey)
	if !verifyUser(user) {
		c.Redirect("/", 307)
		return
	}

	id, _ := c.GetInt("id", 0)
	pTitle := c.GetString("p-title")
	pUrl := c.GetString("p-url")
	solution := c.GetString("solution")
	submissions := c.GetString("submissions")
	rethink := c.GetString("rethink")
	harvest := c.GetString("harvest")
	mark, _ := c.GetBool("mark", false)
	day := c.GetString("day")
	pro := strings.Split(pTitle, ".")
	pName := strings.TrimSpace(strings.Replace(pTitle, pro[0]+".", "", 1))
	pid, err := strconv.Atoi(pro[0])
	logs.Info("id: %d", id)
	if err != nil || pid < 1 || pName == "" {
		c.Ctx.WriteString("ERROR")
		return
	}
	o := orm.NewOrm()
	// 处理题目逻辑
	problem := models.Problem{Pid: pid, Url: pUrl, Name: pName}
	// 三个返回参数依次为：是否新创建的，对象 Id 值，错误
	_, err = o.InsertOrUpdate(&problem, "Pid")
	if err != nil {
		logs.Error(err)
		c.Ctx.WriteString("problem error")
		return
	}
	// 处理日期
	d := time.Now()
	if day != "" {
		if d, err = time.Parse("2006-01-02", day); err != nil {
			d = time.Now()
			logs.Error(err)
		}
	}
	note := models.Note{
		Id:          uint(id),
		Problem:     &problem,
		Language:    models.Language_JAVA,
		Day:         d,
		Solution:    strings.TrimSpace(solution),
		Submissions: strings.TrimSpace(submissions),
		Rethink:     strings.TrimSpace(rethink),
		Harvest:     strings.TrimSpace(harvest),
		Mark:        mark,
	}
	var newId int64
	if id < 1 {
		newId, err = o.Insert(&note)
	} else {
		newId, err = o.Update(&note)
	}
	logs.Info("newId: %d", newId)
	if err != nil {
		logs.Error(err)
		c.Ctx.WriteString("note error")
		return
	}
	c.Ctx.Redirect(http.StatusFound, "/note/"+string(newId))
}
