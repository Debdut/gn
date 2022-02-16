package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MakeNowJust/heredoc"
	"github.com/debdut/gn/pkg/framework/next"
	"github.com/spf13/cobra"
)

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Use:   "page",
	Short: "Generate a Next Page",
	Long: `
		Generate a Next Page with various options
		such as class components, functional components,
		hooks, server side rendering, static generaton.`,
	Example: heredoc.Doc(`
		# create a page interactively
		gn page create
		# create a page, where 'pages' dir is autodetected
		gn page create user/history
		# create a page at a destination
		gn page create about .
	`),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 || len(args) > 3 {
			err := cmd.Help()
			if err != nil {
				panic(err)
			}
			os.Exit(0)
		} else if args[0] != "create" {
			fmt.Printf("%s: no such command \n", args[0])
			os.Exit(0)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			name string
			dir  string
		)
		if len(args) == 2 {
			dir, name = filepath.Split(args[1])
			nextDirs := next.GetNextDirs()
			dir = filepath.Join(nextDirs.Page, dir)
		} else if len(args) == 3 {
			name = args[1]
			dir = args[2]
		}

		err := next.WritePageTemplate(name, dir, false)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
