package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//DB connection global var
var DB *sql.DB

//InitDB allows for creation of DB on bot launch
func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
}
