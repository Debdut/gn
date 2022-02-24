package command

import (
	"fmt"

	"github.com/debdut/gn/pkg/markdown"
)

type CommandExample struct {
	Command string
	Output  string
}

type CommandRun func(*Command, []string) error

type Command struct {
	Name        string
	Command     string
	Usage       string
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

func (c *Command) Help() string {
	md := markdown.Markdown{}

	md.AddHeading(1, c.Command)
	md.AddBlockQuote(c.Short)
	md.AddNewLine()

	md.AddHeading(2, "Usage")
	md.AddParagraph(c.Usage)
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

	return md.Render()
}
