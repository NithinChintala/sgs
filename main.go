package main

import (
	_ "fmt"
	_ "encoding/json"
	_ "database/sql"
	"github.com/NithinChintala/sgs/dao"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"log"
)

var (
	tmpl = template.Must(template.ParseGlob("webapp/*.html"))
)

func main() {
	// Setup the router
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("webapp/static/"))))

	// Setup the API
	api := r.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/papers", dao.GetPapers).Methods("GET")
	api.HandleFunc("/users", dao.GetUsers).Methods("GET")
	api.HandleFunc("/tags", dao.GetTags).Methods("GET")

	// Run the server
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}