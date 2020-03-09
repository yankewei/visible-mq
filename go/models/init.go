package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() (err error) {
	DB, err = sql.Open("mysql", "root:yankewei@/yankewei")
	if err != nil {
		panic(err)
	}
	return
}

func Close() (err error){
	err = DB.Close()
	return
}
