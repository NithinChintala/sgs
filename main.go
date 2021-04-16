package main

import (
	_ "database/sql"
	_ "encoding/json"
	_ "fmt"
	"github.com/NithinChintala/sgs/dao"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var (
	tmpl = template.Must(template.ParseGlob("webapp/*.html"))
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Index", nil)
}

func Papers(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Papers", dao.GetPapers())
}

func PapersByTagWord(w http.ResponseWriter, r *http.Request) {
	tag := r.FormValue("tag")
	log.Println(tag)
	tmpl.ExecuteTemplate(w, "Papers", dao.GetPapersByTagWord(tag))
}

func Users(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Users", dao.GetUsers())
}

func EditPaper(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	id, _ := strconv.Atoi(args["id"])
	paper := dao.GetPapersById(id)

	tmpl.ExecuteTemplate(w, "EditPaper", paper)
}

func main() {
	// Setup the router
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("webapp/static/"))))
	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/papers", Papers).Methods("GET")
	r.Path("/papers/").Queries("tag","{[a-zA-Z0-9]+}").HandlerFunc(PapersByTagWord)
	r.HandleFunc("/papers/{id:[0-9]+}", EditPaper).Methods("GET")
	r.HandleFunc("/users", Users).Methods("GET")

	// Setup the API
	api := r.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/papers", dao.GetPapersAPI).Methods("GET")
	api.HandleFunc("/users", dao.GetUsersAPI).Methods("GET")
	api.HandleFunc("/tags", dao.GetTags).Methods("GET")

	// Run the server
	log.Println("Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
