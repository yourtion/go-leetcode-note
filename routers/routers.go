package routers

import (
	"github.com/astaxie/beego"

	"leetcode-note/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/note/?:id", &controllers.NoteController{})
	beego.Router("/export", &controllers.ExportController{})

}
