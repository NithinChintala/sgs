package dao

import (
	"database/sql"
	"github.com/NithinChintala/sgs/model"
	"log"
	"fmt"
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

func GetPapers() []model.Paper {
	connect()
	defer db.Close()

	results, err := db.Query("SELECT * FROM papers")
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

func CreatePaper(paper model.Paper) {
	connect()
	defer db.Close()

	insert :=
	`
	INSERT INTO papers (year, title, journal, volume, issue, pages, doi)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	_, err := db.Exec(insert, paper.Year, paper.Title, 
		paper.Journal, paper.Volume, paper.Issue, paper.Pages, paper.Doi)
	if err != nil {
		log.Fatal(err)
	}
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

func DeletePaper(id int) {
	connect()
	defer db.Close()

	delete := `
	DELETE FROM papers
	WHERE id=?
	`
	_, err := db.Exec(delete, id)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPapersByUserId(userId int) []model.Paper {
	connect()
	defer db.Close()

	query := `
	SELECT papers.* FROM papers, authors, users
	WHERE papers.id = authors.paper_id
	AND authors.user_id = users.id
	AND users.id = ?
	`

	results, err := db.Query(query, userId)
	if err != nil {
		log.Fatal(err)
	}
	return ReadPapers(results)
}

func GetPapersByCiteeId(citeeId int) []model.Paper {
	connect()
	defer db.Close()

	subQuery := "(SELECT * FROM `references` WHERE citee_id = ?) sub_query"
	query := "SELECT papers.* FROM papers, %s WHERE papers.id = sub_query.citer_id"
	full := fmt.Sprintf(query, subQuery)
	
	results, err := db.Query(full, citeeId)
	if err != nil {
		log.Fatal(err)
	}
	return ReadPapers(results)
}

func GetPapersByCiterId(citerId int) []model.Paper {
	connect()
	defer db.Close()

	subQuery := "(SELECT * FROM `references` WHERE citer_id = ?) sub_query"
	query := "SELECT papers.* FROM papers, %s WHERE papers.id = sub_query.citee_id"
	full := fmt.Sprintf(query, subQuery)
	
	results, err := db.Query(full, citerId)
	if err != nil {
		log.Fatal(err)
	}
	return ReadPapers(results)
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

