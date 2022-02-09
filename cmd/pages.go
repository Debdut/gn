/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pagesCmd represents the pages command
var pagesCmd = &cobra.Command{
	Use:   "pages",
	Short: "Generate a Page",
	Long: `Generate a Page.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pages called")
	},
}

func init() {
	rootCmd.AddCommand(pagesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pagesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pagesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
