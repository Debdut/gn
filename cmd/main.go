package main

import (
	"fmt"
	"os"
	// "github.com/spf13/cobra"
)

var COMMANDS = []string{"api", "page"}

func main() {
	fmt.Printf("%+v\n", os.Args)
}
