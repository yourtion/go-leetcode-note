package models

import "time"

type Language int8

const (
	JAVA Language = iota
	JAVASCRIPT
)

func (l Language) String() string {
	return [...]string{"Java", "JavaScript"}[l]
}

type Note struct {
	Id          uint
	Problem     *Problem  `orm:"rel(fk) "description:"题目ID"`
	Language    Language  `description:"编程语言"`
	Day         time.Time `type(date)description:"日期"`
	Solution    string    `description:"解题思路"`
	Submissions string    `orm:"size(128) "description:"提交记录"`
	Rethink     string    `description:"反思"`
	Harvest     string    `description:"收获"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}

// 多字段索引
func (u *Note) TableIndex() [][]string {
	return [][]string{
		[]string{"Language", "Day"},
	}
}
