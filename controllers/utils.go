package controllers

var LoginKey string
var LoginUser string

func verifyUser(name interface{}) bool {
	return name != nil && name == LoginUser
}
