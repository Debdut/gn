package command

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
