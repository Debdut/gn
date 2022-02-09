package main

import (
	"fmt"
	"os"
	// "github.com/spf13/cobra"
)

var COMMANDS = []string{"api", "page"}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Subcommand missing")
	}

	for i := 0; i < len(COMMANDS); i++ {
		command := COMMANDS[i]
		if os.Args[1] == command {
			fmt.Printf("Subcommand: %s\n", command)
			return
		}
	}

	fmt.Printf("No matching subcommand: %s\n", os.Args[1])
}
