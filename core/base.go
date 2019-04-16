package core

import (
	"github.com/CodeSpecific/douro/kit"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

func init() {
	var err error
	db, err = xorm.NewEngine("mysql", "")
	if err != nil {
		panic(err)
	}
	db.SetMapper(kit.LowerCaseMapper{})
}
