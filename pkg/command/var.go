package command

import "fmt"

type VarCheck func(v *Var, arg string) bool
type VarRun func(v *Var, args []string) error

type Var struct {
	Name string
	Var  string

	Modifiers []*Modifier
	Parent    Arg
	Check     VarCheck
	Run       VarRun

	Short    string // Short Description
	Long     string // Long Description
	Aliases  []string
	Examples []string
}

func (v *Var) Path() []string {
	path := []string{fmt.Sprintf("<%s>", v.Var)}
	if v.Parent != nil {
		path = append(path, v.Parent.Path()...)
	}
	return path
}
