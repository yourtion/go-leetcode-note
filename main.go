package main

import (
	"os"
	"strconv"
	_ "leetcode-note/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

)

func main() {
	logs.SetLogger("console")
	log := logs.GetLogger()

	workingDir := "./"
	// 切换到指定的工作目录
	if err := os.Chdir(workingDir); err != nil {
		logs.Error("change working directory to %s failed: %s", workingDir, err)
	} else {
		logs.Info("current working directory is %s", workingDir)
	}

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
	if os.Getenv("BEEGO_ENV") != "" {
		log.Println("Env $BEEGO_ENV :", os.Getenv("BEEGO_ENV"))
		beego.BConfig.RunMode = os.Getenv("BEEGO_ENV")
	}
	beego.Run()
}
