package controllers

import "github.com/astaxie/beego/logs"

var LoginKey string
var LoginUser string

func verifyUser(name interface{}) bool {
	logs.Trace("verifyUser: %v", name)
	return name != nil && name == LoginUser
}
