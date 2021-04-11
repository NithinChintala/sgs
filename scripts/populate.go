package main

import (
	"database/sql"
	"fmt"
	"github.com/NithinChintala/sgs/dao"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
)

const (
	USER     = "root"
	PASSWORD = "P@ssw0rd"
	HOST     = "localhost"
	PORT     = 3306
	DB_NAME  = "sgs"

	NUM_USERS  = 75
	NUM_PAPERS = 200

	MAX_REFS = 4
	MAX_AUTHS = 4

	INSERT_USER = `
	INSERT INTO users (first_name, last_name, username, password, email)
	VALUES (?, ?, ?, ?, ?)
	`
	INSERT_PAPER = `
	INSERT INTO papers (year, title, volume, issue, pages)
	VALUES (?, ?, ?, ?, ?)
	`
	INSERT_TAG = `INSERT INTO tags (word) VALUES (?)`

	INSERT_AUTHOR    = "INSERT INTO authors (user_id, paper_id) VALUES (?, ?)"
	INSERT_REFERENCE = "INSERT INTO `references` (citer_id, citee_id) VALUES (?, ?)"
	INSERT_KEYWORD   = "INSERT INTO keywords (paper_id, tag_id) VALUES (?, ?)"
)

var (
	db *sql.DB
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Get a random number from [start, end] that is not in seen
func uniqueRandInt(start, end int, seen map[int]bool) int {
	var val int
	for {
		val = (rand.Int() % (end - start)) + start
		if _, ok := seen[val]; !ok {
			break
		}
	}
	return val
}

// Get count unique random numbers from [start, end]
// if count > end - start + 1, then crash 
func uniqueRands(start, end, count int) []int {
	if count > end - start + 1 {
		log.Fatal("Count larger than range")
	}
	seen := make(map[int]bool)
	out := make([]int, 0)
	for i := 0; i < count; i++ {
		val := uniqueRandInt(start, end, seen)
		out = append(out, val)
		seen[val] = true;
	}
	return out
}

// Populate the `users` and `papers` table with random data
func populate() {
	
	for i := 0; i < NUM_USERS; i++ {
		fn := fmt.Sprintf("fn%d", rand.Int()%200)
		ln := fmt.Sprintf("ln%d", rand.Int()%200)
		un := fmt.Sprintf("un%d", rand.Int()%200)
		pw := fmt.Sprintf("pw%d", rand.Int()%200)
		email := fmt.Sprintf("%s@domain%d.edu", un, rand.Int()%200)

		_, err := db.Exec(INSERT_USER, fn, ln, un, pw, email)
		check(err)
	}
	for i := 0; i < NUM_PAPERS; i++ {
		year := (rand.Int() % 61) + 1960
		title := fmt.Sprintf("title%d", rand.Int()%200)
		volume := (rand.Int() % 50) + 1
		issue := (rand.Int() % 12) + 1
		start := (rand.Int() % 1000) + 1
		pages := fmt.Sprintf("%d-%d", start, start+(rand.Int()%20))

		_, err := db.Exec(INSERT_PAPER, year, title, volume, issue, pages)
		check(err)
	}
}

// Create a directed acylic graph (DAG) of papers such that papers 
// published in year n can only reference papers published before n
func createReferences() {
	refMap := make(map[int][]int)
	for i := 0; i < NUM_PAPERS - MAX_REFS; i++ {
		numRefs := rand.Int() % (MAX_REFS + 1)
		refMap[i] = uniqueRands(i + 1, NUM_PAPERS - 1, numRefs)
	}

	results, err := db.Query("SELECT * FROM	papers ORDER BY `year` DESC")
	check(err)

	papers := dao.ReadPapers(results)
	for citer, citees := range refMap {
		citer_id := papers[citer].Id
		for _, i := range citees {
			citee_id := papers[i].Id
			_, err := db.Exec(INSERT_REFERENCE, citer_id, citee_id)
			check(err)
		}
	}
}

func createAuthors() {
}

func main() {

	rand.Seed(1)

	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", USER, PASSWORD, HOST, PORT, DB_NAME)
	var err error
	// using walrus assigns local db variable instead of global
	db, err = sql.Open("mysql", url)
	check(err)
	defer db.Close()

	populate()
	createReferences()
	createAuthors()
}
