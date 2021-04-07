package main

import (
	"fmt"
	"database/sql"
	_ "github.com/NithinChintala/sgs/model"
	_ "github.com/go-sql-driver/mysql"
)

const (
	USER = "root"
	PASSWORD = "P@ssw0rd"
	HOST = "localhost"
	PORT = 3306
	DB_NAME = "sgs"
)

var (
	db *sql.DB
)

func main() {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", USER, PASSWORD, HOST, PORT, DB_NAME)
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}