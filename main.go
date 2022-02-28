package main

import (
	"fmt"

	"github.com/debdut/gn/pkg/cmd/root"
)

func main() {
	rootCmd := root.New()
	fmt.Println(rootCmd.Help())
}
