package command

import (
	"fmt"
	"strings"

	"github.com/debdut/gn/pkg/markdown"
	"github.com/debdut/gn/pkg/util"
)

type CommandRun func(*Command, []string) error

type Command struct {
	Name        string
	Command     string
	Use         string
	Interactive bool

	Commands  []*Command
	Parent    *Command
	Vars      []*Var
	Modifiers []*Modifier
	Run       CommandRun

	Short    string // Short Description
	Long     string // Long Description
	Aliases  []string
	Examples []CommandExample
}

// Returns the path of the command
func (c *Command) Path() []string {
	path := []string{c.Command}
	cmd := c.Parent
	for cmd != nil {
		path = append(path, cmd.Command)
		cmd = cmd.Parent
	}
	util.ReverseStringSlice(path)
	return path
}

// Adds sub commands to this command
func (c *Command) AddCommands(cmds ...*Command) {
	for _, cmd := range cmds {
		if cmd == c {
			panic("Command can't be a child of itself")
		}
		cmd.Parent = c
		c.Commands = append(c.Commands, cmd)
	}
}

// Adds vars to this command
func (c *Command) AddVars(vars ...*Var) {
	for _, v := range vars {
		v.Parent = c
		c.Vars = append(c.Vars, v)
	}
}

// returns usage string
// either from given .Use field
// or constructed from vars
func (c *Command) Usage() string {
	if len(c.Use) > 0 {
		return c.Use
	}

	var text []string
	for _, a := range c.Vars {
		text = append(text, "# "+a.Short)
		text = append(text,
			strings.Join(a.Path(), " ")+"\n")
	}

	return strings.Join(text, "\n")
}

// generate markdown help string
func (c *Command) Help() string {
	md := markdown.Markdown{}

	md.AddHeading(1, c.Command)
	md.AddBlockQuote(c.Short)
	md.AddNewLine()

	usage := c.Usage()
	if len(usage) > 0 {
		md.AddHeading(2, "Usage")
		md.AddCode("sh", c.Usage())
		md.AddNewLine()
	}

	if len(c.Long) > 0 {
		md.AddHeading(2, "Description")
		md.AddParagraph(c.Long)
		md.AddNewLine()
	}

	if len(c.Vars) > 0 {
		headers := []string{"Variables", "Description"}
		var rows [][]string
		for _, v := range c.Vars {
			leaves := v.Leaves()
			for _, leaf := range leaves {
				rows = append(rows, []string{
					leaf.VarPath(),
					leaf.Context,
				})
			}
		}

		md.AddHeading(2, "Variables")
		md.AddTable(headers, rows, []string{})
		md.AddNewLine()
	}

	if len(c.Modifiers) > 0 {
		headers := []string{"Modifier", "Description"}
		var rows [][]string
		for _, mod := range c.Modifiers {
			rows = append(rows, []string{
				mod.Modifier,
				mod.Short,
			})
		}

		md.AddHeading(2, "Modifiers")
		md.AddTable(headers, rows, []string{})
		md.AddNewLine()
	}

	if len(c.Commands) > 0 {
		var subCommands []string
		for _, subCmd := range c.Commands {
			subCommands = append(subCommands,
				fmt.Sprintf("**%s** %s", subCmd.Command, subCmd.Short))
		}

		md.AddHeading(2, "Sub Commands")
		md.AddList(subCommands, [][]string{}, false)
		md.AddNewLine()
	}

	var examples []string
	for _, ex := range c.Examples {
		examples = append(examples, ex.String())
	}

	if len(examples) > 0 {
		md.AddHeading(2, "Examples")
		md.AddCode("sh", strings.Join(examples, "\n"))
		md.AddNewLine()
	}

	return md.Render()
}

func (c *Command) MatchModifiers(arg string) []*Modifier {
	mods := []*Modifier{}
	modNames := ParseModifiers(arg)
	for _, mod := range c.Modifiers {
		for _, modName := range modNames {
			if mod.Modifier == modName {
				mods = append(mods, mod)
			}
		}
	}

	return mods
}

func (c *Command) Match(args []string) []Arg {
	matches := []Arg{c}

	if len(args) == 0 {
		return matches
	}

	arg, rest := args[0], args[1:]
	if strings.HasPrefix(arg, ":") {
		// modifier := arg
		return matches
	}

	matches = []Arg{}
	if len(c.Commands) > 0 {
		for _, cmd := range c.Commands {
			matchF := BeginMatch
			if len(arg) > 2 {
				matchF = HelmMatch
			}
			if matchF(cmd.Command, arg) {
				matches = append(matches, cmd.Match(rest)...)
			}
		}
	}

	if len(c.Vars) > 0 {
		for _, v := range c.Vars {
			if v.Valid(arg) {
				matches = append(matches, v.Match(rest)...)
			}
		}
	}

	return matches
}

func BeginMatch(target string, needle string) bool {
	return strings.HasPrefix(target, needle)
}

func HelmMatch(target string, needle string) bool {
	i, j := 0, 0
	for i < len(target) && j < len(needle) {
		if target[i] == needle[j] {
			j++
		}
		i++
	}

	return j == len(needle)
}
