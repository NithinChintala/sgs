package server

import (
	"html/template"
	"net/http"
	"log"
	"fmt"

	"github.com/gorilla/mux"
)

const (
	host = "localhost"
	port = 8080
	tmplGlob = "webapp/template/*.html"
)

type server struct {
	router *mux.Router
	tmpl *template.Template
}

func (s *server) Run() {
	connStr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Running server at http://%s\n", connStr)
	log.Fatal(http.ListenAndServe(connStr, s.router))
}

func Init() *server {
	log.Println("Initialzing server")
	s := server{}
	s.router = mux.NewRouter()
	s.tmpl = template.Must(template.ParseGlob(tmplGlob))
	s.routes()
	return &s
}

func (s *server) routes() {
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("webapp/static/"))))
	s.router.HandleFunc("/", s.index).Methods("GET")

	s.router.HandleFunc("/papers/{id}/create", s.createPaper).Methods("POST")
	s.router.HandleFunc("/papers", s.getPapers).Methods("GET")
	s.router.HandleFunc("/papers/{id}", s.getPaperById).Methods("GET")
	s.router.HandleFunc("/papers/{id}/update", s.updatePaper).Methods("POST")
	s.router.HandleFunc("/papers/{id}/delete", s.deletePaper).Methods("POST")

	s.router.HandleFunc("/users/{id}/create", s.createUser).Methods("POST")
	s.router.HandleFunc("/users", s.getUsers).Methods("GET")
	s.router.HandleFunc("/users/{id}", s.getUserById).Methods("GET")
	s.router.HandleFunc("/users/{id}/update", s.updateUser).Methods("POST")
	s.router.HandleFunc("/users/{id}/delete", s.deleteUser).Methods("POST")

	s.router.HandleFunc("/tags/{id}/create", s.createTag).Methods("POST")
	s.router.HandleFunc("/tags", s.getTags).Methods("GET")
	s.router.HandleFunc("/tags/{id}", s.getTagById).Methods("GET")
	s.router.HandleFunc("/tags/{id}/update", s.updateTag).Methods("POST")
	s.router.HandleFunc("/tags/{id}/delete", s.deleteTag).Methods("POST")	
}

func (s *server) index(w http.ResponseWriter, r *http.Request) {
	s.tmpl.ExecuteTemplate(w, "Index", nil)
}
