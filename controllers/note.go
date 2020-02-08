package controllers

import (
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
	c.Layout = "layout.html"
	c.TplName = "note/form.tpl"
}

func (c *NoteController) Post() {
	id, _ := c.GetUint8("id")
	pTitle := c.GetString("p-title")
	pUrl := c.GetString("p-url")
	solution := c.GetString("solution")
	submissions := c.GetString("submissions")
	rethink := c.GetString("rethink")
	harvest := c.GetString("harvest")
	mark, _ := c.GetBool("mark", false)
	pro := strings.Split(pTitle, ".")
	pName := strings.TrimSpace(strings.Replace(pTitle, pro[0]+".", "", 1))
	pid, err := strconv.Atoi(pro[0])
	if err != nil || pid < 1 || pName == "" {
		c.Ctx.WriteString("ERROR")
		return
	}
	o := orm.NewOrm()
	// 处理题目逻辑
	problem := models.Problem{Pid: pid, URL: pUrl, Name: pName}
	// 三个返回参数依次为：是否新创建的，对象 Id 值，错误
	_, err = o.InsertOrUpdate(&problem, "Pid")
	if err != nil {
		logs.Error(err)
		c.Ctx.WriteString("problem error")
		return
	}
	note := models.Note{
		Id:          uint(id),
		Problem:     &problem,
		Language:    0,
		Day:         time.Now(),
		Solution:    solution,
		Submissions: submissions,
		Rethink:     rethink,
		Harvest:     harvest,
		Mark:        mark,
	}
	newId, err := o.InsertOrUpdate(&note, "Id")
	if err != nil {
		logs.Error(err)
		c.Ctx.WriteString("note error")
		return
	}
	c.Ctx.Redirect(302, "/note/"+string(newId))
}
