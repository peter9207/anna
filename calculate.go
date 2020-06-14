package main

import "strconv"
import "log"
import "github.com/peter9207/anna/loaders"
import "github.com/spf13/cobra"
import "database/sql"
import _ "github.com/lib/pq"

var insertQuery = ` INSERT INTO results (article_id, cost, tag) VALUES (?, ?, ?) `

var valuesCmd = &cobra.Command{
	Use: "calculate <tag> <filename>",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			cmd.Help()
			return
		}

		tag := args[0]
		filename := args[1]

		var values []float64
		var dates []string

		loaders.FromCSV(filename, func(input []string) {
			f, err := strconv.ParseFloat(input[4], 64)
			if err != nil {
				log.Println(err)
			}
			values = append(values, f)
			dates = append(dates, input[0])
		})

		pois := findPOI(values)
		db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/anna?sslmode=disable")
		if err != nil {
			panic(err)
		}

		for _, v := range pois {

			err = save(dates[v.index], v.value, tag, db)
			if err != nil {
				log.Println(err)
			}
		}

	},
}

var searchQuery = ` SELECT id from articles where tag = ? AND timestamp <= ? order by timestamp desc limit 1 `

func save(date string, cost float64, tag string, db *sql.DB) error {

	var articleId string
	row := db.QueryRow(searchQuery, tag, date)
	err := row.Scan(&articleId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("no article found for %V cost: %v, tag: %v", date, cost, tag)
			return nil
		}
		log.Printf("Error query articles %v", err)
		return err
	}

	_, err = db.Exec(insertQuery, articleId, cost, tag)
	if err != nil {
		log.Printf("Saving results returned err %v ", err)
		return err
	}

	return nil
}
