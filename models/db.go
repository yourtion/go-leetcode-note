package models

import (
	"log"
	"os"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	dataSource := os.Getenv("DATABASE_URL")
	orm.Debug = true

	// set default database
	var err error
	if dataSource != "" {
		err = orm.RegisterDataBase("default", "postgres", dataSource)
		orm.SetMaxOpenConns("default", 20)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_ = orm.RegisterDataBase("default", "sqlite3", "/tmp/leetcode.db")
	}

	orm.RegisterModel(new(Problem), new(Note))
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatal(err)
	}
}
