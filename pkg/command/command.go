package command

type CommandExample struct {
	Command string
	Output  string
}

type CommandRun func(*Command, []string) error

type Command struct {
	Name string

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
func (c *Command) Path() string {
	if c.Parent != nil {
		return c.Parent.Path() + " " + c.Name
	}
	return c.Name
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
