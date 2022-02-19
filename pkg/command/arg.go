package command

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

// Adds args to this command
func (c *Command) AddArgs(args ...*Arg) {
	for _, arg := range args {
		arg.Parent = c
		c.Args = append(c.Args, arg)
	}
}
