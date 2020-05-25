package articles

import (
	"github.com/peter9207/anna/loaders"

	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

type Article struct {
	id          int64
	title       string
	publication string
	author      string
	date        string
	url         string
	content     string
}

var insertQuery = `
INSERT INTO articles (id, title, publication, author, timestamp, url, content)
VALUES($1,$2,$3,$4,$5,$6,$7)
`

// ,id,title,publication,author,date,year,month,url,content
func LoadFromFile(filename string) {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/anna?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Connected to db")

	var count = 0

	loaders.FromCSV(filename, func(values []string) {

		id, err := strconv.Atoi(values[1])
		if err != nil {
			log.Printf("Failed to parse value %v as integer id: %v", values[0], err)
			return
		}

		_, err = db.Exec(insertQuery, id,
			values[2],
			values[3],
			values[4],
			values[5],
			values[8],
			values[9],
		)
		if err != nil {
			log.Print("Error saving record into db", err)
		}

		count = count + 1
		if count%1000 == 0 {
			log.Printf("processed %v records", count)
		}

	})

}
