package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Note struct {
	Id          uint
	Problem     *Problem  `json:"problem"orm:"rel(fk)"description:"题目"`
	Lang        string    `json:"lang"orm:"size(16)"description:"编程语言"`
	Day         time.Time `json:"day"orm:"type(date)"description:"日期"`
	Solution    string    `json:"solution"description:"解题思路"`
	Submissions string    `json:"submissions"orm:"size(128)"description:"提交记录"`
	Rethink     string    `json:"rethink"description:"反思"`
	Harvest     string    `json:"harvest"description:"收获"`
	Mark        bool      `json:"mark"description:"标记"`
	Created     time.Time `json:"created"orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `json:"updated"orm:"auto_now;type(datetime)"`
}

// 多字段索引
func (u *Note) TableIndex() [][]string {
	return [][]string{
		{"Lang", "Day"},
	}
}

func GetNoteById(id int) (error, *Note) {
	note := Note{Id: uint(id)}
	o := orm.NewOrm()
	err := o.Read(&note)
	if err != nil {
		return err, nil
	}
	_, err = o.LoadRelated(&note, "Problem")
	if err != nil {
		return err, nil
	}
	return nil, &note
}
