package main

import (
	"fmt"
	"log"
	"os"

	"github.com/debdut/gn/pkg/cmd/root"
	"github.com/debdut/gn/pkg/command"
)

func main() {
	rootCmd := root.New()
	args := os.Args[1:]
	matches := rootCmd.Match(args)
	if len(matches) == 0 {
		log.Fatalln("command not found")
	} else if len(matches) == 1 {
		cmd := matches[0].(*command.Command)
		fmt.Println(cmd.Help())
	} else {
		fmt.Println("multiples matches found")
		for _, a := range matches {
			cmd := a.(*command.Command)
			fmt.Println(cmd.Path())
		}
	}
}
