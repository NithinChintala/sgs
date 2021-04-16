package dao

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	user = "root"
	password = "P@ssw0rd"
	host = "localhost"
	port = 3306
	db_name = "sgs"
)

var (
	db *sql.DB
	err error
	url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, db_name)
)

// Connect the global db var to the MySQL Server
func connect() {
	db, err = sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}
}