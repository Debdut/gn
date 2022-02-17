package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MakeNowJust/heredoc"
	"github.com/debdut/gn/pkg/framework/next"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Generate a Next Api",
	Long: `
		Generate a Next Api with various methods,
		POST, GET, PUT, DELETE, attach w/ Database
		models.`,
	Example: heredoc.Doc(`
		# create a api interactively
		gn api create
		# create a api, where 'pages' dir is autodetected
		gn api create user/history
		# create a api at a destination
		gn api create about .
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
			dir = filepath.Join(nextDirs.Api, dir)
		} else if len(args) == 3 {
			name = args[1]
			dir = args[2]
		}

		err := next.WriteApiTemplate(name, dir, false)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
