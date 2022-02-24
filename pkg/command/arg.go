package command

import (
	"fmt"
	"strings"
)

type ArgCheck func(*Command, []string) bool
type ArgRun func(*Command, []string) error

type Arg struct {
	Args []string

	Modifiers []string
	Parent    *Command
	Check     ArgCheck
	Run       ArgRun

	Short    string // Short Description
	Long     string // Long Description
	Aliases  []string
	Examples []string
}

func (arg *Arg) Path() []string {
	path := arg.Parent.Path()
	for _, a := range arg.Args {
		path = append(path, fmt.Sprintf("<%s>", a))
	}
	mod := strings.Join(arg.Modifiers, ":")
	path = append(path,
		fmt.Sprintf(":%s:", mod))

	return path
}
