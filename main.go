package main

import (
	"os"
	"strconv"

	"github.com/astaxie/beego/orm"

	"leetcode-note/controllers"
	_ "leetcode-note/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	_ = logs.SetLogger("console")
	log := logs.GetLogger()

	workingDir := "./"
	// 切换到指定的工作目录
	if err := os.Chdir(workingDir); err != nil {
		logs.Error("change working directory to %s failed: %s", workingDir, err)
	} else {
		logs.Info("current working directory is %s", workingDir)
	}
	// 配置端口
	if os.Getenv("PORT") != "" {
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatal("$PORT must be set")
		}
		log.Println("port : ", port)
		beego.BConfig.Listen.HTTPPort = port
	} else {
		beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	}
	// 配置运行环境
	if os.Getenv("BEEGO_ENV") != "" {
		log.Println("Env $BEEGO_ENV :", os.Getenv("BEEGO_ENV"))
		beego.BConfig.RunMode = os.Getenv("BEEGO_ENV")
	}
	// 开启Session
	beego.BConfig.WebConfig.Session.SessionOn = true
	// 配置登录信息
	if os.Getenv("LOGIN_KEY") != "" && os.Getenv("LOGIN_USER") != "" {
		controllers.LoginKey = os.Getenv("LOGIN_KEY")
		controllers.LoginUser = os.Getenv("LOGIN_USER")
		logs.Info("Login info: %s - %s", controllers.LoginKey, controllers.LoginUser)
	}
	// 设置 DEBUG
	if os.Getenv("DEBUG") == "true" {
		logs.Warn("Debug enabled !")
		orm.Debug = true
		logs.SetLevel(logs.LevelTrace)
	} else {
		logs.SetLevel(logs.LevelInfo)
	}
	beego.Run()
}
