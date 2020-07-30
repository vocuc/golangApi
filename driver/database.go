package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //ba la ba la
)

// ConnectSQL ...
func ConnectSQL() (db *sql.DB, err error) {

	dbHost := "localhost"
	dbProt := "3306"
	dbName := "shopeva"
	dbUser := "root"
	dbPass := "123456"

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbProt, dbName)
	db, err = sql.Open("mysql", dbSource)
	return
}
