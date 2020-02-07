package models

import "time"

type Problem struct {
	Pid     int `orm:"pk"`
	URL     string
	Name    string
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}
