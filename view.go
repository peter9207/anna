package main

// import "strconv"
import "log"
import "fmt"
import "text/tabwriter"
import "os"

// import "github.com/peter9207/anna/loaders"
import "github.com/spf13/cobra"
import "database/sql"
import _ "github.com/lib/pq"

var viewCmd = &cobra.Command{

	Use: "view <tag>",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			cmd.Help()
			return
		}

		tag := args[0]

		db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5433/anna?sslmode=disable")
		if err != nil {
			panic(err)
		}

		var query = `

select a.title, a.timestamp, a.author, r.cost
from results r
join articles a on a.id = r.article_id where r.tag = $1
`

		w := tabwriter.NewWriter(os.Stdout, 5, 4, 4, ' ', tabwriter.AlignRight|tabwriter.Debug)

		rows, err := db.Query(query, tag)
		if err != nil {
			log.Printf("Error executing query %v", err)
			return
		}
		defer rows.Close()

		for rows.Next() {

			log.Println("RUNNING")
			var date string
			var title string
			var author string
			var cost float64

			err = rows.Scan(&title, &date, &author, &cost)
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Fprintf(w, "%s\t%s\t%s\t%v\t\n", title, date, author, cost)
		}
		w.Flush()

		return
	},
}
