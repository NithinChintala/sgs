package dao

import (
	"database/sql"
	"github.com/NithinChintala/sgs/model"
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
