package command

import "strings"

type ModRun func(*Modifier) error

type Modifier struct {
	Name     string
	Modifier string

	Parent *Command
	Run    ModRun

	Short    string // Short Description
	Long     string // Long Description
	Aliases  []string
	Examples []string
}

func ParseModifiers(arg string) []string {
	modNames := strings.SplitAfter(arg, ":")
	nonEmptyModNames := []string{}

	for _, m := range modNames {
		if len(m) > 0 {
			nonEmptyModNames = append(nonEmptyModNames, m)
		}
	}

	return nonEmptyModNames
}
