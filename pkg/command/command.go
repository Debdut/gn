package command

import (
	"fmt"
	"strings"

	"github.com/debdut/gn/pkg/markdown"
)

type CommandRun func(*Command, []string) error

type Command struct {
	Name        string
	Command     string
	Use         string
	Interactive bool

	SubCommands [](*Command)
	Parent      *Command
	Args        [](*Arg)
	Modifiers   []string
	Run         CommandRun

	Short    string // Short Description
	Long     string // Long Description
	Aliases  []string
	Examples []CommandExample
}

// Returns the path of the command
func (c *Command) Path() []string {
	path := []string{c.Command}
	if c.Parent != nil {
		path = append(path, c.Parent.Path()...)
	}
	return path
}

// Adds sub commands to this command
func (c *Command) AddCommand(cmds ...*Command) {
	for _, cmd := range cmds {
		if cmd == c {
			panic("Command can't be a child of itself")
		}
		cmd.Parent = c
		c.SubCommands = append(c.SubCommands, cmd)
	}
}

// Adds args to this command
func (c *Command) AddArgs(args ...*Arg) {
	for _, arg := range args {
		arg.Parent = c
		c.Args = append(c.Args, arg)
	}
}

// returns usage string
// either from given .Use field
// or constructed from args
func (c *Command) Usage() string {
	if len(c.Use) > 0 {
		return c.Use
	}

	var text []string
	for _, a := range c.Args {
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

	md.AddHeading(2, "Usage")
	md.AddCode("sh", c.Usage())
	md.AddNewLine()

	md.AddHeading(2, "Description")
	md.AddParagraph(c.Long)
	md.AddNewLine()

	var subCommands []string
	for _, subCmd := range c.SubCommands {
		subCommands = append(subCommands,
			fmt.Sprintf("**%s** %s", subCmd.Command, subCmd.Short))
	}

	md.AddHeading(2, "Sub Commands")
	md.AddList(subCommands, [][]string{}, false)
	md.AddNewLine()

	var examples []string
	for _, ex := range c.Examples {
		examples = append(examples, ex.String())
	}

	md.AddHeading(2, "Examples")
	md.AddCode("sh", strings.Join(examples, "\n"))
	md.AddNewLine()

	return md.Render()
}
