package lib

import (
	"database/sql"
	"time"

	"api/conf"
	_ "github.com/go-sql-driver/mysql"
)

var connect *sql.DB
var dsn string = "root:YclyClycL5819@tcp(" + conf.MYSQLADDRESS + ":3306)/home"

func init() {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	connect = db
}

func GetInstance() *sql.DB {
	//fmt.Println(connect)
	return connect
}
