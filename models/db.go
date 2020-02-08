package models

import (
	"fmt"
	"log"
	"os"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	dataSource := os.Getenv("DATABASE_URL")
	orm.Debug = true

	// set default database
	var err error
	if dataSource != "" {
		fmt.Printf(dataSource)
		err = orm.RegisterDataBase("default", "postgres", dataSource)
		orm.SetMaxOpenConns("default", 20)
		if err != nil {
			log.Fatal(err)
		}
	}

	orm.RegisterModel(new(Problem), new(Note))
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatal(err)
	}
}
