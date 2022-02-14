/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (

	// "github.com/debdut/gn/pkg/cmd"
	"fmt"

	"github.com/debdut/gn/pkg/framework/next"
)

func main() {
	// cmd.Execute()
	dirs := next.GetNextDirs()
	fmt.Printf("%v\n", dirs)

	ts := next.IsTypescript()
	fmt.Println(ts)

	confs := next.GetConfigs()
	fmt.Printf("%v\n", confs)
}
