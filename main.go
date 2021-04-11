package main

import (
	"fmt"
	"encoding/json"
	"database/sql"
	"github.com/NithinChintala/sgs/dao"
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

	result, err := db.Query("SELECT * FROM papers")
	papers := dao.ReadPapers(result)
	b, err := json.Marshal(papers)
	fmt.Println(string(b))
}