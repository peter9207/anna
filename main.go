package main

// import "strconv"
// import "log"
import "github.com/spf13/cobra"
import "fmt"
import "github.com/peter9207/anna/articles"

var rootCmd = &cobra.Command{
	Use: "anna",
	// Short: "my attempt at doing some data analytics",
	// Long:  "Too lazy to write this",
}

var articlesCmd = &cobra.Command{
	Use:   "articles <function>",
	Short: "articles functions",
}

var loadArticlesCmd = &cobra.Command{
	Use:   "load filename",
	Short: "load articles from a csv file",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			cmd.Help()
			return
		}

		articles.LoadFromFile(args[0])

	},
}

func main() {

	articlesCmd.AddCommand(loadArticlesCmd)

	rootCmd.AddCommand(articlesCmd)

	rootCmd.Execute()

	// var values []float64

	// var dates []string

	// file := "data/GSPC.csv"
	// var cb = func(input []string) {
	// 	f, err := strconv.ParseFloat(input[4], 64)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	values = append(values, f)
	// 	dates = append(dates, input[0])
	// }
	// ReadCSV(file, cb)

	// results := findPOI(values)

	// log.Printf("Found %v results\n", len(results))
}
