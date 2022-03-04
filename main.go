package main

import (
	"fmt"

	"github.com/debdut/gn/pkg/cmd/root"
	"github.com/debdut/gn/pkg/command"
)

func main() {
	rootCmd := root.New()
	// fmt.Println(rootCmd.Help())

	matches := rootCmd.Match([]string{"next", "page", "create"})
	for _, a := range matches {
		cmd := a.(*command.Command)
		fmt.Printf("%s\n", cmd.Path())
	}
}
