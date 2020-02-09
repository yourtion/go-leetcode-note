package controllers

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"
)

var LoginKey string
var LoginUser string

func verifyUser(name interface{}) bool {
	logs.Trace("verifyUser: %v", name)
	return name != nil && name == LoginUser
}

func getWeekRange() (weekMonday int, weekSunday int) {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday, _ = strconv.Atoi(weekStartDate.Format("20060102"))
	weekSunday, _ = strconv.Atoi(weekStartDate.AddDate(0, 0, 6).Format("20060102"))
	return
}
