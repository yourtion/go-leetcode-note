package models

import "time"

type Problem struct {
	Pid     int       `json:"pid"orm:"pk"description:"题目ID"`
	Url     string    `json:"url"description:"题目URL"`
	Name    string    `json:"name"description:"题目名称"`
	Created time.Time `json:"created"orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated"orm:"auto_now;type(datetime)"`
}
