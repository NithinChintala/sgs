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

func UpdatePaper(id int, paper model.Paper) {
	connect()
	defer db.Close()

	update := `
	UPDATE papers
	SET id=?, year=?, title=?, journal=?, volume=?, issue=?, pages=?, doi=?
	WHERE id=?
	`

	_, err := db.Exec(update, paper.Id, paper.Year, paper.Title, 
		paper.Journal, paper.Volume, paper.Issue, paper.Pages, paper.Doi,
		paper.Id)
	if err != nil {
		log.Fatal(err)
	}
}

func CreatePaper(paper model.Paper) {

}

func DeletePaper(id int) {

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

func GetPapersByTagWord(tag string) []model.Paper {
	connect()
	defer db.Close()

	query := `
	SELECT papers.* FROM papers, keywords, tags
	WHERE papers.id = keywords.paper_id 
	AND keywords.tag_id = tags.id 
	AND tags.word =?
	`

	results, err := db.Query(query, tag)
	if err != nil {
		log.Fatal(err)
	}
	return ReadPapers(results)
}

func GetPapersById(id int) []model.Paper {
	connect()
	defer db.Close()

	results, err := db.Query("SELECT * FROM papers WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	return ReadPapers(results)
}

func GetPapers() []model.Paper {
	connect()
	defer db.Close()

	results, err := db.Query("SELECT * FROM papers")
	if err != nil {
		log.Fatal(err)
	}
	return ReadPapers(results)
}

func GetPapersAPI(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(GetPapers())
}
