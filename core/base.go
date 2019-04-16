package core

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

func init() {
	var err error
	db, err = xorm.NewEngine("mysql", "****:****@(*****:*****)/****?charset=utf8")
	if err != nil {
		panic(err)
	}
}
