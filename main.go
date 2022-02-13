/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (

	// "github.com/debdut/gn/pkg/cmd"
	"fmt"

	"github.com/debdut/gn/pkg/dir"
)

func main() {
	// cmd.Execute()
	dirs := dir.GetNextDirs()
	fmt.Printf("%v\n", dirs)
}
