package dao

import (
	"database/sql"
	"github.com/NithinChintala/sgs/model"
	"net/http"
	"encoding/json"
	"log"
)

func ReadPapers(result *sql.Rows) []model.Paper {
	papers := make([]model.Paper, 0)
	for result.Next() {
		var paper model.Paper
		err := result.Scan(&paper.Id, &paper.Year, &paper.Title,
			&paper.Journal, &paper.Volume, &paper.Issue,
			&paper.Pages, &paper.Doi)
		if err != nil {
			log.Fatal(err)
		}
		papers = append(papers, paper)
	}

	return papers
}

func GetPapersByUserId(w http.ResponseWriter, r *http.Request) {
	connect()
	defer db.Close()

	query := `
	SELECT papers.* FROM papers, authors, users
	WHERE papers.id = authors.paper_id
	AND authors.user_id = users.id
	AND users.id = ?
	`

	results, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	papers := ReadPapers(results)
	json.NewEncoder(w).Encode(papers)
}

func GetPapersByTagWord(w http.ResponseWriter, r *http.Request) {
	connect()
	defer db.Close()

	query := `
	SELECT papers.* FROM papers, keywords, tags
	WHERE papers.id = keywords.paper_id 
	AND keywords.tag_id = tags.id 
	AND tags.word = ?
	`

	results, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	papers := ReadPapers(results)
	json.NewEncoder(w).Encode(papers)
}

func GetPapers(w http.ResponseWriter, r *http.Request) {
	connect()
	defer db.Close()

	results, err := db.Query("SELECT * FROM papers")
	if err != nil {
		log.Fatal(err)
	}
	papers := ReadPapers(results)
	json.NewEncoder(w).Encode(papers)
}
