package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//DB connection global var
var (
	DB     *sql.DB
	config string
)

//Config for DB setup
type Config struct {
	DataSourceName string `json:"dataSourceName"`
}

//InitDB allows for creation of DB on bot launch
func InitDB() {
	var err error
	DB, err = sql.Open("mysql", config)
	if err != nil {
		panic(err.Error())
	}
	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
}

//LoadConfig to load DB connectstring
func LoadConfig() {
	conf, e := ioutil.ReadFile("./conf.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var c []Config
	jsonErr := json.Unmarshal(conf, &c)
	if jsonErr != nil {
		log.Panic(jsonErr)
	}

	config = c[0].DataSourceName

}
