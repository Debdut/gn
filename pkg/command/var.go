package command

import (
	"fmt"
	"sort"
	"strings"

	"github.com/debdut/gn/pkg/util"
)

type VarCheck func(v *Var, arg string) bool
type VarRun func(v *Var, args []string) error

type Var struct {
	Name string
	Var  string

	Vars []*Var
	// Modifiers []*Modifier
	Parent Arg
	Check  VarCheck
	Run    VarRun

	Short    string // Short Description
	Long     string // Long Description
	Context  string
	Aliases  []string
	Examples []string
}

func (v *Var) Path() []string {
	path := v.pathHelper()
	util.ReverseStringSlice(path)

	return path
}

func (v *Var) pathHelper() []string {
	path := []string{fmt.Sprintf("<%s>", v.Var)}
	if v.Parent != nil {
		path = append(path, v.Parent.Path()...)
	}
	return path
}

func isVar(v interface{}) bool {
	switch v.(type) {
	case Var:
		return true
	default:
		return false
	}
}

func (v *Var) VarPathTrail() []*Var {
	path := []*Var{v}
	node := v.Parent
	for node != nil && isVar(node) {
		nv := (node).(*Var)
		path = append(path, nv)
		node = nv.Parent
	}

	// reverse
	sort.Slice(path, func(i, j int) bool { return i > j })

	return path
}

func (v *Var) VarPath() string {
	vs := v.VarPathTrail()
	path := []string{}
	for _, v := range vs {
		path = append(path, fmt.Sprintf("<%s>", v.Var))
	}

	return strings.Join(path, " ")
}

func (v *Var) Leaves() []*Var {
	var (
		next   []*Var
		leaves []*Var
	)
	v.LeavesHelper(next, leaves)

	return leaves
}

func (v *Var) LeavesHelper(next []*Var, leaves []*Var) {
	for _, vi := range v.Vars {
		if len(vi.Vars) == 0 {
			leaves = append(leaves, vi)
		} else {
			next = append(next, vi)
			vi.LeavesHelper(next, leaves)
		}
	}
}

func (v *Var) Valid(arg string) bool {
	return v.Check(v, arg)
}

func (v *Var) Match(args []string) []Arg {
	matches := []Arg{v}
	if len(args) == 0 {
		return matches
	}

	arg, rest := args[0], args[1:]
	if strings.HasPrefix(arg, ":") {
		// modifier := arg
		return matches
	}

	if len(v.Vars) > 0 {
		for _, v := range v.Vars {
			if v.Valid(arg) {
				matches = append(matches, v.Match(rest)...)
			}
		}
	}

	return matches
}
