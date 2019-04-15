package core

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

func init() {
	var err error
	db, err = xorm.NewEngine("mysql", "root:123456@(106.12.130.92:3306)/douro?charset=utf8")
	if err != nil {
		panic(err)
	}
}
