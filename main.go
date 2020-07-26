package main

// import "strconv"
// import "log"
import "github.com/spf13/cobra"

// import "github.com/peter9207/anna/loaders"

// import "fmt"
import "github.com/peter9207/anna/articles"

var rootCmd = &cobra.Command{
	Use: "anna",
}

var articlesCmd = &cobra.Command{
	Use:   "articles <function>",
	Short: "articles functions",
}

var loadArticlesCmd = &cobra.Command{
	Use:   "load filename",
	Short: "load articles from a csv file",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 2 {
			cmd.Help()
			return
		}

		articles.LoadFromFile(args[0], args[1])
	},
}

var determineCmd = &cobra.Command{
	Use:   "determine <function>",
	Short: "determine functions",
}

func main() {

	articlesCmd.AddCommand(loadArticlesCmd)
	determineCmd.AddCommand(expCoeff)
	rootCmd.AddCommand(articlesCmd)
	rootCmd.AddCommand(determineCmd)
	rootCmd.AddCommand(viewCmd)
	rootCmd.AddCommand(valuesCmd)
	rootCmd.Execute()

}
