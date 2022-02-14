/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	_ "github.com/debdut/gn/pkg/cmd"
	"github.com/debdut/gn/pkg/framework/next"

	// tmplNext "github.com/debdut/gn/pkg/template/next"
	"github.com/debdut/gn/pkg/util"
)

func main() {
	// cmd.Execute()
	// dirs := next.GetNextDirs()
	// fmt.Printf("%v\n", dirs)

	// ts := next.IsTypescript()
	// fmt.Println(ts)

	// confs := next.GetConfigs()
	// fmt.Printf("%v\n", confs)

	// api := tmplNext.Api{Api: "Invoice", TS: true}
	// err := tmplNext.GenApi(api, os.Stdout)
	// if err != nil {
	// 	panic(err)
	// }

	err := next.WritePageTemplate("about", util.Cwd(), true)
	if err != nil {
		panic(err)
	}
}
