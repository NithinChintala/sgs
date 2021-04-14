package main

import (
	"fmt"
	"encoding/json"
	"database/sql"
	"github.com/NithinChintala/sgs/dao"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"log"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "webapp/" + r.URL.Path[1:])
	})

	http.HandleFunc("/api/papers", func(w http.ResponseWriter, r *http.Request) {
		results, err := db.Query("SELECT * FROM papers")
		if err != nil {
			log.Fatal(err)
		}
		papers := dao.ReadPapers(results)
		json.NewEncoder(w).Encode(papers)
	})

	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		results, err := db.Query("SELECT * FROM users")
		if err != nil {
			log.Fatal(err)
		}
		users := dao.ReadUsers(results)
		json.NewEncoder(w).Encode(users)
	})

	http.HandleFunc("/api/tags", func(w http.ResponseWriter, r *http.Request) {
		results, err := db.Query("SELECT * FROM tags")
		if err != nil {
			log.Fatal(err)
		}
		tags := dao.ReadTags(results)
		json.NewEncoder(w).Encode(tags)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}