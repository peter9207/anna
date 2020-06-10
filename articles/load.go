package articles

import (
	"github.com/peter9207/anna/loaders"

	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"strings"
)

type Article struct {
	id          int64
	title       string
	publication string
	author      string
	date        string
	url         string
	content     string
	tag         string
}

var insertQuery = `
INSERT INTO articles (id, title, publication, author, timestamp, url, content, tag)
VALUES($1,$2,$3,$4,$5,$6,$7, $8)
`

func searchIgnoreCase(s, v string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(v))
}

func LoadFromFile(filename string, tag string) {
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

	var total = 0
	var count = 0

	loaders.FromCSV(filename, func(values []string) {

		id, err := strconv.Atoi(values[1])
		if err != nil {
			log.Printf("Failed to parse value %v as integer id: %v", values[0], err)
			return
		}

		title := values[2]
		content := values[9]

		inTitle := searchIgnoreCase(title, tag)
		inContents := searchIgnoreCase(content, tag)

		if inTitle || inContents {
			_, err = db.Exec(insertQuery, id,
				values[2],
				values[3],
				values[4],
				values[5],
				values[8],
				values[9],
				tag,
			)
			if err != nil {
				log.Print("Error saving record into db", err)
			}
			count = count + 1
		}

		total = total + 1
		if count%1000 == 0 {
			log.Printf("processed %v records %v inserted", total, count)
		}

	})

}
